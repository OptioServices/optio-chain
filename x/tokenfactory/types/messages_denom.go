package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDenom{}

func NewMsgCreateDenom(
	owner string,
	denom string,
	description string,
	ticker string,
	precision int32,
	url string,
	maxSupply int64,
	supply int64,
	canChangeMaxSupply bool,
	limitDailyMinting bool,
	dailyMintingLimit int64,
	lastMintDate string,
	hasYearlyHalving bool,
	nextHalvingDate string,

) *MsgCreateDenom {
	return &MsgCreateDenom{
		Owner:              owner,
		Denom:              denom,
		Description:        description,
		Ticker:             ticker,
		Precision:          precision,
		Url:                url,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
		LimitDailyMinting:  limitDailyMinting,
		DailyMintingLimit:  dailyMintingLimit,
		HasYearlyHalving:   hasYearlyHalving,
	}
}

func (msg *MsgCreateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	tickerLength := len(msg.Ticker)
	if tickerLength < 3 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Ticker length must be at least 3 chars long")
	}
	if tickerLength > 10 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Ticker length must be 10 chars long maximum")
	}
	if msg.MaxSupply == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Max Supply must be greater than 0")
	}
	if msg.LimitDailyMinting && msg.DailyMintingLimit == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Daily Minting Limit must be greater than 0")
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDenom{}

func NewMsgUpdateDenom(
	owner string,
	denom string,
	description string,
	url string,
	maxSupply int64,
	supply int32,
	canChangeMaxSupply bool,

) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Owner:              owner,
		Denom:              denom,
		Description:        description,
		Url:                url,
		MaxSupply:          maxSupply,
		CanChangeMaxSupply: canChangeMaxSupply,
	}
}

func (msg *MsgUpdateDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid owner address (%s)", err)
	}
	if msg.MaxSupply == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "Max Supply must be greater than 0")
	}
	return nil
}
