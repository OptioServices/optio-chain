package keeper

import (
	"context"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	"github.com/OptioServices/optio/x/distribute/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) Distribute(goCtx context.Context, msg *types.MsgDistribute) (*types.MsgDistributeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	authorized := k.IsAuthorized(ctx, msg.Creator)
	if !authorized {
		return nil, errorsmod.Wrapf(sdkerrors.ErrUnauthorized, "Unauthorized sender")
	}

	currentSupply := k.bankKeeper.GetSupply(ctx, msg.Amount.Denom).Amount.Uint64()
	if currentSupply+msg.Amount.Amount.Uint64() > k.GetParams(ctx).MaxSupply {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Max supply exceeded")
	}

	err := k.bankKeeper.MintCoins(ctx, types.ModuleName, sdk.NewCoins(msg.Amount))
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Minting coins failed")
	}

	for _, recipient := range msg.Recipients {
		acct, err := sdk.AccAddressFromBech32(recipient.Address)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid recipient address")
		}

		coins := sdk.NewCoins(sdk.NewCoin(msg.Amount.Denom, math.NewIntFromUint64(recipient.Amount)))
		err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, acct, coins)
		if err != nil {
			return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Sending coins failed")
		}
	}

	return &types.MsgDistributeResponse{}, nil
}
