package keeper

import (
	"context"
	"time"

	"github.com/charleenfei/modules/incubator/faucet/internal/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) QueryMining(c context.Context, req *types.QueryMiningRequest) (*types.QueryMiningResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	a, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mintTime := ctx.BlockTime().Unix()
	m := k.getMining(ctx, a.String())
	isPresent := k.isPresent(ctx, m.Minter)
	var timeLeft int64
	if !isPresent {
		timeLeft = 0
	} else {
		lastTime := time.Unix(m.LastTime, 0)
		currentTime := time.Unix(mintTime, 0)

		lastTimePlusLimit := lastTime.Add(k.Limit).UTC()
		isAfter := lastTimePlusLimit.After(currentTime)
		if isAfter {
			timeLeft = int64(lastTime.Add(k.Limit).UTC().Sub(currentTime).Seconds())
		} else {
			timeLeft = 0
		}
	}

	return &types.QueryMiningResponse{
		TimeLeft: timeLeft,
	}, nil
}

func (k Keeper) QueryFaucet(c context.Context, req *types.QueryFaucetRequest) (*types.QueryFaucetResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	value := k.GetFaucetKey(ctx)
	return &types.QueryFaucetResponse{
		Key: value,
	}, nil
}

// query mining
// query faucet
