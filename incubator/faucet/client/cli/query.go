package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"

	// "github.com/cosmos/cosmos-sdk/client/context"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/charleenfei/modules/incubator/faucet/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group pooltoy queries under a subcommand
	pooltoyQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	pooltoyQueryCmd.AddCommand(
	// flags.GetCommands(
	// 	GetCmdWhenBrrr(queryRoute, cdc),
	// )...,
	)

	return pooltoyQueryCmd
}

// func GetCmdWhenBrrr(queryRoute string, cdc *codec.Marshaler) *cobra.Command {
// 	return &cobra.Command{
// 		Use:   "when-brrr [userAccount]",
// 		Short: "how many seconds until this user can brrr again",
// 		Args:  cobra.ExactArgs(1),
// 		RunE: func(cmd *cobra.Command, args []string) error {
// 			ctx, err := client.GetClientQueryContext(cmd)
// 			if err != nil {
// 				return err
// 			}
// 			address := args[0]
// 			_, err = sdk.AccAddressFromBech32(address)
// 			if err != nil {
// 				fmt.Printf("could not query User\n%s\n", err.Error())
// 				return nil
// 			}
// 			res, _, err := ctx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, keeper.QueryWhenBrrr, address), nil)
// 			if err != nil {
// 				fmt.Printf("could not query User\n%s\n", err.Error())
// 				return nil
// 			}
// 			var out int64
// 			cdc.MustUnmarshalJSON(res, &out)
// 			return ctx.PrintOutput(out)
// 		},
// 	}
// }
