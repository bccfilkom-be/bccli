package cmd

import (
	"os"

	"github.com/bccfilkom-be/bccli/cmd/domain"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "cmd",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(domain.DomainCmd)
}
