package keeper

import (
	"context"
	"time"

	"optio/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateDenom(goCtx context.Context, msg *types.MsgCreateDenom) (*types.MsgCreateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value already exists
	_, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "index already set")
	}

	dailyMintingLimit := msg.DailyMintingLimit
	if msg.LimitDailyMinting && msg.HasYearlyHalving {
		now := time.Now()
		if (isLeapYear(now.Year()) && (now.Month() <= 2 && now.Day() <= 29)) || (isLeapYear(now.AddDate(1, 0, 0).Year()) && (now.Month() > 2)) {
			dailyMintingLimit = msg.MaxSupply / 2 / 366
		} else {
			dailyMintingLimit = msg.MaxSupply / 2 / 365
		}
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Ticker:             msg.Ticker,
		Precision:          msg.Precision,
		Url:                msg.Url,
		MaxSupply:          msg.MaxSupply,
		CanChangeMaxSupply: msg.CanChangeMaxSupply,
		LimitDailyMinting:  msg.LimitDailyMinting,
		DailyMintingLimit:  dailyMintingLimit,
		MintedToday:        0,
		LastMintDate:       "1970-01-01",
		HasYearlyHalving:   msg.HasYearlyHalving,
		NextHalvingDate:    time.Now().Format("1970-01-01"),
		DateCreated:        time.Now().Format("1970-01-01"),
		Supply:             0,
	}

	if !msg.LimitDailyMinting {
		denom.DailyMintingLimit = 0
	}

	k.SetDenom(
		ctx,
		denom,
	)
	return &types.MsgCreateDenomResponse{}, nil
}

func (k msgServer) UpdateDenom(goCtx context.Context, msg *types.MsgUpdateDenom) (*types.MsgUpdateDenomResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Check if the value exists
	valFound, isFound := k.GetDenom(
		ctx,
		msg.Denom,
	)
	if !isFound {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, "Denom not found")
	}

	// Checks if the msg owner is the same as the current owner
	if msg.Owner != valFound.Owner {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	maxSupply, canChangeMaxSupply := valFound.MaxSupply, valFound.CanChangeMaxSupply
	if valFound.CanChangeMaxSupply {
		maxSupply = msg.MaxSupply
		canChangeMaxSupply = msg.CanChangeMaxSupply
	}

	var denom = types.Denom{
		Owner:              msg.Owner,
		Denom:              msg.Denom,
		Description:        msg.Description,
		Url:                msg.Url,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}

	k.SetDenom(ctx, denom)

	return &types.MsgUpdateDenomResponse{}, nil
}
