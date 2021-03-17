package keeper

import (
	"context"

	"github.com/charleenfei/modules/incubator/faucet/internal/types"
)

var _ types.MsgServer = Keeper{}

func (k Keeper) Mint(c context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	return &types.MsgMintResponse{}, nil
}

func (k Keeper) Mining(c context.Context, msg *types.MsgMining) (*types.MsgMiningResponse, error) {
	return &types.MsgMiningResponse{}, nil
}

func (k Keeper) FaucetKey(c context.Context, msg *types.MsgFaucetKey) (*types.MsgFaucetKeyResponse, error) {
	return &types.MsgFaucetKeyResponse{}, nil
}

// func handleMsgMint(ctx sdk.Context, keeper Keeper, msg types.MsgMint) (*sdk.Result, error) {
// 	keeper.Logger(ctx).Info("received mint message: %s", msg)

// 	results := emoji.FindAll(msg.Denom)
// 	if len(results) != 1 {
// 		return nil, types.ErrNoEmoji
// 	}
// 	emo, ok := results[0].Match.(emoji.Emoji)
// 	if !ok {
// 		return nil, types.ErrNoEmoji
// 	}
// 	msg.Denom = emo.Value

// 	time := ctx.BlockTime().Unix()
// 	err := keeper.MintAndSend(ctx, msg.Sender, msg.Minter, time, msg.Denom)
// 	if err != nil {
// 		fmt.Println("err", err)
// 		return nil, sdkerrors.Wrap(err, fmt.Sprintf(" in [%v] hours", keeper.Limit.Hours()))
// 	}

// 	return &sdk.Result{}, nil // return
// }

// // Handle a message to Mint
// func handleMsgFaucetKey(ctx sdk.Context, keeper Keeper, msg types.MsgFaucetKey) (*sdk.Result, error) {

// 	keeper.Logger(ctx).Info("received faucet message: %s", msg)
// 	if keeper.HasFaucetKey(ctx) {
// 		return nil, types.ErrFaucetKeyExisted
// 	}

// 	keeper.SetFaucetKey(ctx, msg.Armor)

// 	return &sdk.Result{}, nil // return
// }
