package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgDistribute{}

func NewMsgDistribute(creator string, amount uint64, recipients []Output) *MsgDistribute {
	return &MsgDistribute{
		FromAddress: creator,
		Amount:      amount,
		Outputs:     recipients,
	}
}

func (msg *MsgDistribute) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.FromAddress)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}

	if msg.Amount == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount cannot be zero")
	}

	if len(msg.Outputs) == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "recipients cannot be empty")
	}

	var total uint64
	for _, recipient := range msg.Outputs {
		_, err := sdk.AccAddressFromBech32(recipient.Address)
		if err != nil {
			return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address (%s)", err)
		}
		total += recipient.Amount
	}
	if total != msg.Amount {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "amount and recipients total do not match")
	}
	return nil
}
