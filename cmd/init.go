package cmd

import (
	"github.com/bccfilkom-be/bccli/internal/framework"
	"github.com/bccfilkom-be/bccli/internal/gocmd"
	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

const (
	successCreateFile      string = "successed: Make file %s\n"
	successCreateDirectory string = "successed: Make Directory %s\n"
)

var Framework string

func init() {
	initCmd.Flags().StringVar(&Framework, "framework", "chi", "web framework of choice, [chi]")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "generate go project for building REST API",
	Long: `
This command is used for make project with name that you specified.
This project purposed for building REST API with clean architecture inspired by Uncle Bob`,
	Args: cobra.ExactArgs(1),
	RunE: _init,
}

func _init(cmd *cobra.Command, args []string) error {
	arg := stringy.New(args[0])
	name := arg.SnakeCase().ToLower()

	if err := gocmd.Init(name); err != nil {
		return err
	}
	if err := framework.Main(framework.NewFramework(Framework)); err != nil {
		return err
	}

	return nil
}
