package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"lto/x/lto/types"
)

var _ = strconv.Itoa(0)

func CmdCreateAnchor() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-anchor [hash] [encoding]",
		Short: "Broadcast message createAnchor",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argHash := args[0]
			argEncoding := args[1]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAnchor(
				clientCtx.GetFromAddress().String(),
				argHash,
				argEncoding,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
