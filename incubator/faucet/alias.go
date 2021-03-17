package faucet

import (
	"github.com/charleenfei/modules/incubator/faucet/internal/keeper"
	"github.com/charleenfei/modules/incubator/faucet/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	MAINNET = "mainnet"
	TESTNET = "testnet"
)

var (
	NewKeeper     = keeper.NewKeeper
	ModuleCdc     = types.ModuleCdc
	RegisterCodec = types.RegisterCodec
)

type (
	Keeper = keeper.Keeper
)
