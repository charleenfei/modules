package keeper

import (
	"fmt"
	"time"

	"github.com/charleenfei/modules/incubator/faucet/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	auth "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	"github.com/tendermint/tendermint/libs/log"
	emoji "github.com/tmdvs/Go-Emoji-Utils"
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
func (k Keeper) MintAndSend(ctx sdk.Context, msg *types.MsgMint) error {
	results := emoji.FindAll(msg.Denom)
	if len(results) != 1 {
		return types.ErrNoEmoji
	}

	emo, ok := results[0].Match.(emoji.Emoji)
	if !ok {
		return types.ErrNoEmoji
	}

	msg.Denom = emo.Value
	mintTime := ctx.BlockTime().Unix()
	if msg.Denom == k.StakingKeeper.BondDenom(ctx) {
		return types.ErrCantWithdrawStake
	}

	// refuse mint in 24 hours
	mining := k.getMining(ctx, msg.Sender)
	if k.isPresent(ctx, msg.Sender) &&
		time.Unix(int64(mining.Lasttime), 0).Add(k.Limit).UTC().After(time.Unix(mintTime, 0)) {
		return types.ErrWithdrawTooOften
	}

	newCoin := sdk.NewCoin(msg.Denom, sdk.NewInt(k.amount))
	mining.Tally = mining.Tally + k.amount
	mining.Lasttime = mintTime
	k.setMining(ctx, msg.Sender, mining)
	k.Logger(ctx).Info("Mint coin: %s", newCoin)
	newCoins := sdk.NewCoins(newCoin)
	if err := k.SupplyKeeper.MintCoins(ctx, types.ModuleName, newCoins); err != nil {
		return err
	}

	r, err := sdk.AccAddressFromBech32(msg.Minter)
	if err != nil {
		return err
	}

	receiverAccount := k.AccountKeeper.GetAccount(ctx, r)
	if receiverAccount == nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s does not exist and is not allowed to receive tokens", msg.Minter)
	}

	if err := k.SupplyKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, r, sdk.NewCoins(newCoin)); err != nil {
		return err
	}

	return nil
}

func (k Keeper) getMining(ctx sdk.Context, minter string) types.MsgMining {
	store := ctx.KVStore(k.storeKey)
	if !k.isPresent(ctx, minter) {
		return *types.NewMining(minter, 0)
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
