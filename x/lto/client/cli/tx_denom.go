package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"lto-cosmos/x/lto/types"
)

func CmdCreateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-denom [denom] [description] [ticket] [precision] [url] [max-supply] [supply] [can-change-max-supply]",
		Short: "Create a new Denom",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexDenom := args[0]

			// Get value arguments
			argDescription := args[1]
			argTicket := args[2]
			argPrecision, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argUrl := args[4]
			argMaxSupply, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}
			argSupply, err := cast.ToInt32E(args[6])
			if err != nil {
				return err
			}
			argCanChangeMaxSupply, err := cast.ToBoolE(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDenom(
				clientCtx.GetFromAddress().String(),
				indexDenom,
				argDescription,
				argTicket,
				argPrecision,
				argUrl,
				argMaxSupply,
				argSupply,
				argCanChangeMaxSupply,
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

func CmdUpdateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-denom [denom] [description] [ticket] [precision] [url] [max-supply] [supply] [can-change-max-supply]",
		Short: "Update a Denom",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			// Get indexes
			indexDenom := args[0]

			// Get value arguments
			argDescription := args[1]
			argTicket := args[2]
			argPrecision, err := cast.ToInt32E(args[3])
			if err != nil {
				return err
			}
			argUrl := args[4]
			argMaxSupply, err := cast.ToInt32E(args[5])
			if err != nil {
				return err
			}
			argSupply, err := cast.ToInt32E(args[6])
			if err != nil {
				return err
			}
			argCanChangeMaxSupply, err := cast.ToBoolE(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDenom(
				clientCtx.GetFromAddress().String(),
				indexDenom,
				argDescription,
				argTicket,
				argPrecision,
				argUrl,
				argMaxSupply,
				argSupply,
				argCanChangeMaxSupply,
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

func CmdDeleteDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-denom [denom]",
		Short: "Delete a Denom",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			indexDenom := args[0]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteDenom(
				clientCtx.GetFromAddress().String(),
				indexDenom,
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
