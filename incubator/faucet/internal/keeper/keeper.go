package keeper

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/charleenfei/modules/incubator/faucet/internal/types"
)

const FaucetStoreKey = "DefaultFaucetStoreKey"

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	SupplyKeeper  types.SupplyKeeper
	StakingKeeper types.StakingKeeper
	AccountKeeper auth.AccountKeeper
	amount        int64                 // set default amount for each mint.
	Limit         time.Duration         // rate limiting for mint, etc 24 * time.Hours
	storeKey      sdk.StoreKey          // Unexposed key to access store from sdk.Context
	cdc           codec.BinaryMarshaler //
}

// NewKeeper creates new instances of the Faucet Keeper
func NewKeeper(
	supplyKeeper types.SupplyKeeper,
	stakingKeeper types.StakingKeeper,
	accountKeeper auth.AccountKeeper,
	amount int64,
	rateLimit time.Duration,
	storeKey sdk.StoreKey,
	cdc codec.BinaryMarshaler) Keeper {
	return Keeper{
		SupplyKeeper:  supplyKeeper,
		StakingKeeper: stakingKeeper,
		AccountKeeper: accountKeeper,
		amount:        amount,
		Limit:         rateLimit,
		storeKey:      storeKey,
		cdc:           cdc,
	}
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// MintAndSend mint coins and send to receiver.
func (k Keeper) MintAndSend(ctx sdk.Context, sender string, receiver string, mintTime int64, denom string) error {

	if denom == k.StakingKeeper.BondDenom(ctx) {
		return types.ErrCantWithdrawStake
	}

	mining := k.getMining(ctx, sender)
	// refuse mint in 24 hours
	if k.isPresent(ctx, sender) &&
		time.Unix(int64(mining.Lasttime), 0).Add(k.Limit).UTC().After(time.Unix(mintTime, 0)) {
		return types.ErrWithdrawTooOften
	}
	newCoin := sdk.NewCoin(denom, sdk.NewInt(k.amount))
	mining.Tally = mining.Tally + k.amount
	mining.Lasttime = mintTime
	k.setMining(ctx, sender, mining)
	k.Logger(ctx).Info("Mint coin: %s", newCoin)
	newCoins := sdk.NewCoins(newCoin)
	err := k.SupplyKeeper.MintCoins(ctx, types.ModuleName, newCoins)
	if err != nil {
		return err
	}

	r, err := sdk.AccAddressFromBech32(receiver)
	if err != nil {
		return err
	}
	receiverAccount := k.AccountKeeper.GetAccount(ctx, r)
	if receiverAccount == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s does not exist and is not allowed to receive tokens", receiver)
	}

	err = k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, r, sdk.NewCoins(newCoin))
	if err != nil {
		return err
	}
	return nil
}

func (k Keeper) getMining(ctx sdk.Context, minter string) types.MsgMining {
	store := ctx.KVStore(k.storeKey)
	if !k.isPresent(ctx, minter) {
		return types.NewMining(minter, 0)
	}
	bz := store.Get([]byte(minter))
	var mining types.MsgMining
	k.cdc.MustUnmarshalBinaryBare(bz, &mining)
	return mining
}

func (k Keeper) setMining(ctx sdk.Context, minter string, mining types.MsgMining) {
	if mining.Minter == "" {
		return
	}
	if mining.Tally == 0 {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(minter), k.cdc.MustMarshalBinaryBare(&mining))
}

// IsPresent check if the name is present in the store or not
func (k Keeper) isPresent(ctx sdk.Context, minter string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(minter))
}

func (k Keeper) GetFaucetKey(ctx sdk.Context) types.MsgFaucetKey {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get([]byte(FaucetStoreKey))
	var faucet types.MsgFaucetKey
	k.cdc.MustUnmarshalBinaryBare(bz, &faucet)
	return faucet
}

func (k Keeper) SetFaucetKey(ctx sdk.Context, armor string) {
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(FaucetStoreKey), []byte(armor))
}

func (k Keeper) HasFaucetKey(ctx sdk.Context) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(FaucetStoreKey))
}
