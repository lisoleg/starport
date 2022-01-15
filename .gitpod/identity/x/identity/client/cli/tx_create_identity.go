package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/lisoleg/identity/x/identity/types"
)

var _ = strconv.Itoa(0)

func CmdCreateIdentity() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-identity [id] [pub-key] [certificate] [credentials] [owner] [data]",
		Short: "Broadcast message CreateIdentity",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argId := args[0]
			argPubKey := args[1]
			argCertificate := args[2]
			argCredentials := args[3]
			argOwner := args[4]
			argData := args[5]

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateIdentity(
				clientCtx.GetFromAddress().String(),
				argId,
				argPubKey,
				argCertificate,
				argCredentials,
				argOwner,
				argData,
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
