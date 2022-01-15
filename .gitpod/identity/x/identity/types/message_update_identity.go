package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateIdentity{}

func NewMsgUpdateIdentity(creator string, id string, pubKey string, certificate string, credentials string, owner string, data string) *MsgUpdateIdentity {
	return &MsgUpdateIdentity{
		Creator:     creator,
		Id:          id,
		PubKey:      pubKey,
		Certificate: certificate,
		Credentials: credentials,
		Owner:       owner,
		Data:        data,
	}
}

func (msg *MsgUpdateIdentity) Route() string {
	return RouterKey
}

func (msg *MsgUpdateIdentity) Type() string {
	return "UpdateIdentity"
}

func (msg *MsgUpdateIdentity) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateIdentity) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateIdentity) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
