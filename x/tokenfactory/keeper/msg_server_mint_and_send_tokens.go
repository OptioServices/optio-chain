package keeper

import (
	"context"
	gomath "math"
	"time"

	"optio/x/tokenfactory/types"

	errorsmod "cosmossdk.io/errors"
	math "cosmossdk.io/math"
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

	now := time.Now()
	if denom.HasYearlyHalving {
		const layout = "2020-12-24"
		nextHalvingdate, err := time.Parse(layout, denom.NextHalvingDate)
		if err != nil {
			panic("Error parsing the denom's nextHalvingDate")
		}

		if now.Year() == nextHalvingdate.Year() && now.Month() == nextHalvingdate.Month() && now.Day() == nextHalvingdate.Day() {
			timesHalved := denom.TimesHalved + 1
			daysInYear := 365
			if (isLeapYear(now.Year()) && (now.Month() <= 2 && now.Day() <= 29)) || (isLeapYear(now.AddDate(1, 0, 0).Year()) && (now.Month() > 2)) {
				daysInYear = 366
			}

			dailyMintingLimit := denom.MaxSupply / int64(gomath.Pow(float64(2), float64(timesHalved+1))) / int64(daysInYear)

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
				DailyMintingLimit:  dailyMintingLimit,
				MintedToday:        denom.MintedToday,
				LastMintDate:       denom.LastMintDate,
				HasYearlyHalving:   denom.HasYearlyHalving,
				NextHalvingDate:    time.Now().Format("2024-12-24"),
			}
			k.SetDenom(ctx, updateDenom)
		}
	}

	var mintedToday int64 = 0
	if denom.LimitDailyMinting {
		const layout = "2020-12-24"
		lastMintdate, err := time.Parse(layout, denom.LastMintDate)
		if err != nil {
			panic("Error parsing the denom's lastMintDate")
		}

		if now.Year() == lastMintdate.Year() && now.Month() == lastMintdate.Month() && now.Day() == lastMintdate.Day() {
			if denom.DailyMintingLimit < denom.DailyMintingLimit+msg.Amount {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "minted amount exceeds daily minting limit")
			}
			mintedToday = denom.MintedToday + msg.Amount
		} else {
			if denom.DailyMintingLimit < msg.Amount {
				return nil, errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "minted amount exceeds daily minting limit")
			}
			mintedToday = msg.Amount

		}
	}

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
		MintedToday:        mintedToday,
		LastMintDate:       time.Now().Format("2020-12-24"),
		HasYearlyHalving:   denom.HasYearlyHalving,
	}
	k.SetDenom(ctx, updateDenom)

	return &types.MsgMintAndSendTokensResponse{}, nil
}

func isLeapYear(year int) bool {
	// A leap year occurs on every year that is evenly divisible by 4
	// except for century years (years ending with 00)
	// which must be divisible by 400.
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		} else {
			return true
		}
	} else {
		return false
	}
}
