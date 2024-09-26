package domain

import (
	"github.com/bccfilkom-be/bccli/cmd"
	"github.com/spf13/cobra"
)

var domainCmd = &cobra.Command{
	Use: "domain",
	Short: "Generate domain components like handler, usecase, and repository.",
}

func init() {
	cmd.RootCmd.AddCommand(domainCmd)
}
