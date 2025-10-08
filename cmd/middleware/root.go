package middleware

import (
	"github.com/bccfilkom-be/bccli/cmd"
	"github.com/spf13/cobra"
)

var middlewareCmd = &cobra.Command{
	Use: "middleware <command",
	Short: "Middleware layer command",
}

func init() {
	cmd.RootCmd.AddCommand(middlewareCmd)
}