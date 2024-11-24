package types

import (
	"testing"

	"github.com/OptioServices/optio-chain/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgDistribute_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDistribute
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDistribute{
				FromAddress: "invalid_address",
				Amount:      100,
				Outputs: []Output{
					{
						Address: sample.AccAddress(),
						Amount:  100,
					},
				},
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDistribute{
				FromAddress: sample.AccAddress(),
				Amount:      100,
				Outputs: []Output{
					{
						Address: sample.AccAddress(),
						Amount:  100,
					},
				},
			},
		}, {
			name: "mismatched amount",
			msg: MsgDistribute{
				FromAddress: sample.AccAddress(),
				Amount:      100,
				Outputs: []Output{
					{
						Address: sample.AccAddress(),
						Amount:  100,
					},
					{
						Address: sample.AccAddress(),
						Amount:  100,
					},
				},
			},
			err: sdkerrors.ErrInvalidRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
