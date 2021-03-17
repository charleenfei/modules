package cli

import (
	"fmt"

	"github.com/charleenfei/modules/incubator/faucet/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
)

// GetTxCmd return faucet sub-command for tx
func GetTxCmd() *cobra.Command {
	faucetTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "faucet transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	faucetTxCmd.AddCommand(
		GetCmdMint(),
		GetCmdMintFor(),
	)

	return faucetTxCmd
}

// GetCmdMint is the CLI command for mining coin
func GetCmdMint() *cobra.Command {
	return &cobra.Command{
		Use:   "mint",
		Short: "mint coin to sender address",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			denom := args[0]
			minter := ctx.GetFromAddress()
			msg := types.NewMsgMint(minter, minter, denom)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}
}

// GetCmdMintFor is the CLI command for mining coin
func GetCmdMintFor() *cobra.Command {
	return &cobra.Command{
		Use:   "mintfor [address] [denom]",
		Short: "mint coin for new address",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			minter, _ := sdk.AccAddressFromBech32(args[0])
			sender := ctx.GetFromAddress()
			denom := args[1]
			msg := types.NewMsgMint(sender, minter, denom)
			if err = msg.ValidateBasic(); err != nil {
				return fmt.Errorf("message validation failed: %w", err)
			}

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}
}
