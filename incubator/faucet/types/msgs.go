package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

var (
	_ sdk.Msg = &MsgMint{}
)

const (
	TypeMint = "mint"
)

// NewMsgMint is a constructor function for NewMsgMint
func NewMsgMint(sender sdk.AccAddress, minter sdk.AccAddress, denom string) *MsgMint {
	return &MsgMint{Sender: sender.String(), Minter: minter.String(), Denom: denom}
}

// Route should return the name of the module
func (msg *MsgMint) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgMint) Type() string { return TypeMint }

// ValidateBasic runs stateless checks on the message
func (msg *MsgMint) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Minter)
	}
	_, err = sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Minter)
	}
	results := emoji.FindAll(msg.Denom)
	if len(results) != 1 {
		return ErrNoEmoji
	}

	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgMint) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *MsgMint) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Sender)}
}
