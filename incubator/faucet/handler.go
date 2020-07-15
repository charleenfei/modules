package faucet

import (
	"fmt"

	"github.com/okwme/modules/incubator/faucet/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
)

// NewHandler returns a handler for "faucet" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		switch msg := msg.(type) {
		case types.MsgMint:
			return handleMsgMint(ctx, keeper, msg)
		case types.MsgFaucetKey:
			return handleMsgFaucetKey(ctx, keeper, msg)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, fmt.Sprintf("Unrecognized faucet Msg type: %v", msg.Type()))
		}
	}
}

// Handle a message to Mint
func handleMsgMint(ctx sdk.Context, keeper Keeper, msg types.MsgMint) (*sdk.Result, error) {
	keeper.Logger(ctx).Info("received mint message: %s", msg)

	results := emoji.FindAll(msg.Denom)
	if len(results) != 1 {
		return nil, types.ErrNoEmoji
	}
	emo, ok := results[0].Match.(emoji.Emoji)
	if !ok {
		return nil, types.ErrNoEmoji
	}
	msg.Denom = emo.Value

	err := keeper.MintAndSend(ctx, msg.Minter, msg.Time, msg.Denom)
	if err != nil {
		fmt.Println("err", err)
		return nil, sdkerrors.Wrap(err, fmt.Sprintf(" in [%v] hours", keeper.Limit.Hours()))
	}

	return &sdk.Result{}, nil // return
}

// Handle a message to Mint
func handleMsgFaucetKey(ctx sdk.Context, keeper Keeper, msg types.MsgFaucetKey) (*sdk.Result, error) {

	keeper.Logger(ctx).Info("received faucet message: %s", msg)
	if keeper.HasFaucetKey(ctx) {
		return nil, types.ErrFaucetKeyExisted
	}

	keeper.SetFaucetKey(ctx, msg.Armor)

	return &sdk.Result{}, nil // return
}
