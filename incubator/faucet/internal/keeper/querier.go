package keeper

import (
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// query endpoints supported by the nameservice Querier
const (
	QueryFaucetKey = "key"
	QueryWhenBrrr  = "whenBrr"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err error) {
		switch path[0] {
		case QueryWhenBrrr:
			return queryWhenBrrr(ctx, path[1:], req, keeper)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown faucet query endpoint")
		}
	}
}

func queryWhenBrrr(ctx sdk.Context, path []string, req abci.RequestQuery, k Keeper) ([]byte, error) {
	accountAddress := path[0]
	userAccount, err := sdk.AccAddressFromBech32(accountAddress)
	if err != nil {
		return nil, err
	}
	mintTime := ctx.BlockTime().Unix()
	mining := k.getMining(ctx, userAccount)
	var timeLeft int64
	isPresent := k.isPresent(ctx, mining.Minter)
	if !isPresent {
		timeLeft = 0
	} else {
		lastTime := time.Unix(mining.LastTime, 0)
		currentTime := time.Unix(mintTime, 0)

		lastTimePlusLimit := lastTime.Add(k.Limit).UTC()
		isAfter := lastTimePlusLimit.After(currentTime)
		if isAfter {
			timeLeft = int64(lastTime.Add(k.Limit).UTC().Sub(currentTime).Seconds())
		} else {
			timeLeft = 0
		}
	}

	res, err := codec.MarshalJSONIndent(k.cdc, timeLeft)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}
	return res, nil
}

func queryFaucetKey(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	value := keeper.GetFaucetKey(ctx)
	res, err := codec.MarshalJSONIndent(keeper.cdc, value)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
