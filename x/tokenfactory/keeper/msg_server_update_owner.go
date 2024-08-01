package keeper

import (
	"context"

	"optio/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateOwner(goCtx context.Context, msg *types.MsgUpdateOwner) (*types.MsgUpdateOwnerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	denom, isFound := k.GetDenom(ctx, msg.Denom)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "denom not found")
	}

	if denom.Owner != msg.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "invalid owner")
	}

	var updateDenom = types.Denom{
		Owner:              msg.NewOwner,
		Denom:              msg.Denom,
		Description:        denom.Description,
		MaxSupply:          denom.MaxSupply,
		Supply:             denom.Supply,
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

	return &types.MsgUpdateOwnerResponse{}, nil
}
