package cmd

import (
	"github.com/spf13/cobra"
	"github.com/sumup-oss/go-pkgs/os"
)

const rootCmdName = "payments-bank-account-golang-assignment"

func NewRootCmd(osExecutor os.OsExecutor) *cobra.Command {
	cmdInstance := &cobra.Command{
		Use:           rootCmdName,
		Short:         rootCmdName,
		Long:          rootCmdName,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	cmdInstance.AddCommand(
		NewApiCmd(osExecutor),
	)

	return cmdInstance
}
