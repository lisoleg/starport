package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmonaut/planet/x/blog/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreatePost())
	cmd.AddCommand(CmdUpdatePost())
	cmd.AddCommand(CmdDeletePost())
	cmd.AddCommand(CmdCreateSentPost())
	cmd.AddCommand(CmdUpdateSentPost())
	cmd.AddCommand(CmdDeleteSentPost())
	cmd.AddCommand(CmdCreateTimedoutPost())
	cmd.AddCommand(CmdUpdateTimedoutPost())
	cmd.AddCommand(CmdDeleteTimedoutPost())
	cmd.AddCommand(CmdSendIbcPost())
	// this line is used by starport scaffolding # 1

	return cmd
}
