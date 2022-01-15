package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateIdentity{}

func NewMsgCreateIdentity(creator string, id string, pubKey string, certificate string, credentials string, owner string, data string) *MsgCreateIdentity {
	return &MsgCreateIdentity{
		Creator:     creator,
		Id:          id,
		PubKey:      pubKey,
		Certificate: certificate,
		Credentials: credentials,
		Owner:       owner,
		Data:        data,
	}
}

func (msg *MsgCreateIdentity) Route() string {
	return RouterKey
}

func (msg *MsgCreateIdentity) Type() string {
	return "CreateIdentity"
}

func (msg *MsgCreateIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
