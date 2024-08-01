package keeper

import (
	"context"

	"optio/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) MintAndSendTokens(goCtx context.Context, msg *types.MsgMintAndSendTokens) (*types.MsgMintAndSendTokensResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom, isFound := k.GetDenom(ctx, msg.Denom)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom not found")
	}

	if denom.Owner != msg.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "denom owner is not the same as the sender")
	}

	if denom.MaxSupply < denom.Supply+msg.Amount {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "minted amount exceeds max supply")
	}

	// TODO: Implement the daily limit check. How will we store/update what has been minted for the day?
	// Maybe just one call will mint all tokens, send them to the distributor address, which will then
	// distribute the tokens? Don't love the transferring twice.

	moduleAccount := k.accountKeeper.GetModuleAddress(types.ModuleName)
	recipientAddress, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	var mintCoins sdk.Coins

	mintCoins = mintCoins.Add(sdk.NewCoin(msg.Denom, math.NewInt(int64(msg.Amount))))
	if err := k.bankKeeper.MintCoins(ctx, types.ModuleName, mintCoins); err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SendCoins(ctx, moduleAccount, recipientAddress, mintCoins); err != nil {
		return nil, err
	}

	var updateDenom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        denom.Description,
		MaxSupply:          denom.MaxSupply,
		Supply:             denom.Supply + msg.Amount,
		Precision:          denom.Precision,
		Ticker:             denom.Ticker,
		Url:                denom.Url,
		CanChangeMaxSupply: denom.CanChangeMaxSupply,
		LimitDailyMinting:  denom.LimitDailyMinting,
		DailyMintingLimit:  denom.DailyMintingLimit,
		HasHalving:         denom.HasHalving,
		YearsToHalving:     denom.YearsToHalving,
	}
	k.SetDenom(ctx, updateDenom)

	return &types.MsgMintAndSendTokensResponse{}, nil
}
