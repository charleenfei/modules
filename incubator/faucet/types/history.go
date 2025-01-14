package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMintHistory = "mint-history"
)

// NewMining returns a new Mining Message
func NewMintHistory(minter sdk.AccAddress, tally int64) *MintHistory {
	return &MintHistory{
		Minter:   minter.String(),
		Lasttime: 0,
		Tally:    tally,
	}
}

// Route should return the name of the module
func (msg MintHistory) Route() string { return RouterKey }

// Type should return the action
func (msg MintHistory) Type() string { return TypeMintHistory }

// ValidateBasic runs stateless checks on the message
func (msg MintHistory) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, msg.Minter)
	}
	return nil
}

// GetSigners defines whose signature is required
func (msg MintHistory) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}
