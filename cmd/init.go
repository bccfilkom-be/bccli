package cmd

import (
	"runtime"

	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/framework"
	"github.com/bccfilkom-be/bccli/internal/gocmd"
	"github.com/bccfilkom-be/bccli/internal/template"
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
	Short: "Initialize a new Go REST server project structure.",
	Long:  "Bootstraps a new Go REST server project by generating the required files\nand directories for running a simple server.",
	Args:  cobra.ExactArgs(1),
	RunE:  _init,
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
	dockerfile, err := file.Create("Dockerfile")
	if err != nil {
		return err
	}
	if err := template.Execute(dockerfile, "Dockerfile", map[string]interface{}{"GoVersion": runtime.Version()[2:]}); err != nil {
		return err
	}
	dockerignore, err := file.Create(".dockerignore")
	if err != nil {
		return err
	}
	if err := template.Execute(dockerignore, "dockerignore", nil); err != nil {
		return err
	}
	makefile, err := file.Create("Makefile")
	if err != nil {
		return err
	}
	if err := template.Execute(makefile, "Makefile", nil); err != nil {
		return err
	}

	return nil
}
