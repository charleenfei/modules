package keeper

import (
	"context"
	"fmt"

	"github.com/charleenfei/modules/incubator/faucet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type MsgServer struct {
	k Keeper
}

func (m MsgServer) Mint(c context.Context, msg *types.MsgMint) (*types.MsgMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	if err := m.k.MintAndSend(ctx, msg); err != nil {
		// TODO: does this error still make sense?
		return nil, sdkerrors.Wrap(err, fmt.Sprintf(" in [%v] hours", m.k.Limit.Hours()))
	}
	return &types.MsgMintResponse{}, nil
}
