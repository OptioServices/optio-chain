package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDistribute{}

func NewMsgDistribute(creator string, amount sdk.Coin, recipients []*MsgDistribute_Output) *MsgDistribute {
	return &MsgDistribute{
		Creator:    creator,
		Amount:     amount,
		Recipients: recipients,
	}
}

func (msg *MsgDistribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Amount.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount cannot be zero")
	}

	if len(msg.Recipients) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "recipients cannot be empty")
	}

	var total uint64
	for _, recipient := range msg.Recipients {
		_, err := sdk.AccAddressFromBech32(recipient.Address)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
		}
		total += recipient.Amount
	}
	if total != msg.Amount.Amount.Uint64() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount and recipients total do not match")
	}
	return nil
}
