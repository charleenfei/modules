package keeper

import (
	"context"
	"time"

	"github.com/charleenfei/modules/incubator/faucet/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type QueryServer struct {
	k Keeper
}

func (q QueryServer) QueryWhenBrr(c context.Context, req *types.QueryWhenBrrRequest) (*types.QueryWhenBrrResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	a, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	mintTime := ctx.BlockTime().Unix()
	m := q.k.getMintHistory(ctx, a)
	ma, err := sdk.AccAddressFromBech32(m.Minter)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	isPresent := q.k.isPresent(ctx, ma)
	var timeLeft int64
	if !isPresent {
		return &types.QueryWhenBrrResponse{
			TimeLeft: 0,
		}, nil
	}

	lastTime := time.Unix(m.Lasttime, 0)
	currentTime := time.Unix(mintTime, 0)

	lastTimePlusLimit := lastTime.Add(q.k.Limit).UTC()
	isAfter := lastTimePlusLimit.After(currentTime)
	if isAfter {
		timeLeft = int64(lastTime.Add(q.k.Limit).UTC().Sub(currentTime).Seconds())
	} else {
		timeLeft = 0
	}

	return &types.QueryWhenBrrResponse{
		TimeLeft: timeLeft,
	}, nil
}
