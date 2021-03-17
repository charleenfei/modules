package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

var (
	_ sdk.Msg = &MsgMint{}
	_ sdk.Msg = &MsgFaucetKey{}
	_ sdk.Msg = &MsgMining{}
)

const (
	TypeMint      = "mint"
	TypeFaucetKey = "faucet_key"
	TypeMining    = "mining"
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
	_, err := sdk.ValAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Minter)
	}
	_, err = sdk.ValAddressFromBech32(msg.Sender)
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

// FAUCET KEY
// NewMsgFaucetKey is a constructor function for MsgFaucetKey
func NewMsgFaucetKey(sender sdk.AccAddress, armor string) MsgFaucetKey {
	return MsgFaucetKey{Sender: sender.String(), Armor: armor}
}

// Route should return the name of the module
func (msg *MsgFaucetKey) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgFaucetKey) Type() string { return TypeFaucetKey }

// ValidateBasic runs stateless checks on the message
func (msg *MsgFaucetKey) ValidateBasic() error {
	if len(msg.Armor) == 0 {
		return ErrFaucetKeyEmpty
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgFaucetKey) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *MsgFaucetKey) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

// MINING
// NewMining returns a new Mining
func NewMining(minter string, tally int64) *MsgMining {
	return &MsgMining{
		Minter:   minter,
		Lasttime: 0,
		Tally:    tally,
	}
}

// Route should return the name of the module
func (msg *MsgMining) Route() string { return RouterKey }

// Type should return the action
func (msg *MsgMining) Type() string { return TypeMining }

// ValidateBasic runs stateless checks on the message
func (msg *MsgMining) ValidateBasic() error {
	_, err := sdk.ValAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Minter)
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg *MsgMining) GetSignBytes() []byte {
	panic("amino support disabled")
}

// GetSigners defines whose signature is required
func (msg *MsgMining) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
