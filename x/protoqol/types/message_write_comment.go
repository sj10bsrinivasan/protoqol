package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgWriteComment = "write_comment"

var _ sdk.Msg = &MsgWriteComment{}

func NewMsgWriteComment(creator string) *MsgWriteComment {
	return &MsgWriteComment{
		Creator: creator,
	}
}

func (msg *MsgWriteComment) Route() string {
	return RouterKey
}

func (msg *MsgWriteComment) Type() string {
	return TypeMsgWriteComment
}

func (msg *MsgWriteComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgWriteComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgWriteComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
