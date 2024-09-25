package infra

import (
	"github.com/bccfilkom-be/bccli/cmd"
	"github.com/spf13/cobra"
)

var infraCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new Go REST server project structure.",
	Long:  "Bootstraps a new Go REST server project by generating the required files\nand directories for running a simple server.",
	Args:  cobra.ExactArgs(1),
}

func init() {
	cmd.RootCmd.AddCommand(infraCmd)
}
