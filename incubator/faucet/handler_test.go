package faucet

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/kyokomi/emoji"
	"github.com/stretchr/testify/require"

	"github.com/okwme/modules/incubator/faucet/internal/types"
	"github.com/tendermint/tendermint/crypto"
)

func TestEmoji(t *testing.T) {
	moduleAcct := sdk.AccAddress(crypto.AddressHash([]byte("foo")))
	moduleAcct2 := sdk.AccAddress(crypto.AddressHash([]byte("bar")))
	denom := "🥵"
	msg := types.NewMsgMint(moduleAcct, moduleAcct2, time.Now().Unix(), denom)

	err := msg.ValidateBasic()
	require.NoError(t, err)

	msg.Denom = emoji.Sprint(msg.Denom)
	codeWords := emoji.RevCodeMap()[msg.Denom]

	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		require.NoError(t, err)
	}

	if len(codeWords) > 0 {
		msg.Denom = reg.ReplaceAllString(codeWords[0], "")
		require.True(t, true)
	} else {
		fmt.Println("failed to find emoji in msg.Denom", msg.Denom)
		require.True(t, false)
	}
	fmt.Println("final msg.Denom", msg.Denom)
	// require.True(t, false)
}
