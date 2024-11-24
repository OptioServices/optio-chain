package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/OptioServices/optio-chain/x/distribute/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Distribute(goCtx context.Context, msg *types.MsgDistribute) (*types.MsgDistributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorized := k.IsAuthorized(ctx, msg.FromAddress)
	if !authorized {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Unauthorized sender")
	}

	denom := k.GetParams(ctx).Denom
	currentSupply := k.bankKeeper.GetSupply(ctx, denom).Amount.Uint64()
	if currentSupply+msg.Amount > k.GetParams(ctx).MaxSupply {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Max supply exceeded")
	}

	coin := sdk.NewCoin(denom, math.NewIntFromUint64(msg.Amount))
	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(coin))
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Minting coins failed")
	}

	for _, recipient := range msg.Outputs {
		acct, err := sdk.AccAddressFromBech32(recipient.Address)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid recipient address")
		}

		coins := sdk.NewCoins(sdk.NewCoin(denom, recipient.Coin.Amount))
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, acct, coins)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Sending coins failed")
		}
	}

	return &types.MsgDistributeResponse{}, nil
}
