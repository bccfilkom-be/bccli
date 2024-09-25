package cmd

import (
	"os"

	"github.com/bccfilkom-be/bccli/cmd/domain"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "cmd",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(domain.DomainCmd)
}
