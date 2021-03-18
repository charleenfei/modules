package faucet

import (
	"github.com/charleenfei/modules/incubator/faucet/keeper"
	"github.com/charleenfei/modules/incubator/faucet/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey

	MAINNET = "mainnet"
	TESTNET = "testnet"
)

var (
	NewKeeper = keeper.NewKeeper
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper = keeper.Keeper
)
