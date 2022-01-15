package types

import (
	"testing"

	"github.com/cosmonaut/planet/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateTimedoutPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateTimedoutPost
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateTimedoutPost{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateTimedoutPost{
				Creator: sample.AccAddress(),
			},
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

func TestMsgUpdateTimedoutPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateTimedoutPost
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateTimedoutPost{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateTimedoutPost{
				Creator: sample.AccAddress(),
			},
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

func TestMsgDeleteTimedoutPost_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteTimedoutPost
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteTimedoutPost{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteTimedoutPost{
				Creator: sample.AccAddress(),
			},
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
