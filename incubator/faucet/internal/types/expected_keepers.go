package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SupplyKeeper is required for mining coin
type SupplyKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(
		ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins,
	) error
	GetSupply(ctx sdk.Context, denom string) sdk.Coin
}

// StakingKeeper is required for getting Denom
type StakingKeeper interface {
	BondDenom(ctx sdk.Context) string
}
