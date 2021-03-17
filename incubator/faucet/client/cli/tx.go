package cli

import (
	"bufio"
	"errors"
	"fmt"

	"github.com/charleenfei/modules/incubator/faucet/internal/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
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
		GetCmdInitial(),
	)

	return faucetTxCmd
}

// GetCmdWithdraw is the CLI command for mining coin
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

// GetCmdWithdraw is the CLI command for mining coin
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

// func GetPublishKey(cdc *codec.Marshaler) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "publish",
// 		Short: "Publish current account as an public faucet. Do NOT add many coins in this account",
// 		Args:  cobra.ExactArgs(0),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			ctx, err := client.GetClientTxContext(cmd)
// 			if err != nil {
// 				return err
// 			}

// 			inBuf := bufio.NewReader(cmd.InOrStdin())
// 			home, err := cmd.Flags().GetString("home")
// 			if err != nil {
// 				return err
// 			}

// 			backend, err := cmd.Flags().GetString("keyring-backend")
// 			if err != nil {
// 				return err
// 			}

// 			kb, err := keyring.New(sdk.KeyringServiceName(), backend, home, inBuf)
// 			if err != nil {
// 				return err
// 			}

// 			// check local key
// 			sender := ctx.GetFromAddress()
// 			armor, err := kb.ExportPubKeyArmorByAddress(sender)
// 			if err != nil {
// 				return err
// 			}

// 			msg := types.NewMsgFaucetKey(sender, armor)
// 			err = msg.ValidateBasic()
// 			if err != nil {
// 				return err
// 			}

// 			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
// 		},
// 	}
// }

func GetCmdInitial() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize mint key for faucet",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			inBuf := bufio.NewReader(cmd.InOrStdin())
			home, err := cmd.Flags().GetString("home")
			if err != nil {
				return err
			}

			backend, err := cmd.Flags().GetString("keyring-backend")
			if err != nil {
				return err
			}

			kb, err := keyring.New(sdk.KeyringServiceName(), backend, home, inBuf)
			if err != nil {
				return err
			}

			// check local key
			_, err = kb.Key(types.ModuleName)
			if err == nil {
				return errors.New("faucet existed")
			}

			armor, err := kb.ExportPubKeyArmor(types.ModuleName)
			if err != nil {
				return nil
			}
			kb.ImportPubKey(types.ModuleName, armor)
			return nil
		},
	}
}
