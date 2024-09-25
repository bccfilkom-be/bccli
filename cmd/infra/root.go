package infra

import (
	"github.com/bccfilkom-be/bccli/cmd"
	"github.com/spf13/cobra"
)

var infraCmd = &cobra.Command{
	Use:   "infra <command>",
	Short: "Infra layer command",
}

func init() {
	cmd.RootCmd.AddCommand(infraCmd)
}
