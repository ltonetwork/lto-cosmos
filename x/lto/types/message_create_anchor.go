package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateAnchor = "create_anchor"

var _ sdk.Msg = &MsgCreateAnchor{}

func NewMsgCreateAnchor(creator string, hash string, encoding string) *MsgCreateAnchor {
	return &MsgCreateAnchor{
		Creator:  creator,
		Hash:     hash,
		Encoding: encoding,
	}
}

func (msg *MsgCreateAnchor) Route() string {
	return RouterKey
}

func (msg *MsgCreateAnchor) Type() string {
	return TypeMsgCreateAnchor
}

func (msg *MsgCreateAnchor) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAnchor) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAnchor) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
