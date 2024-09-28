package cmd

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/framework"
	"github.com/bccfilkom-be/bccli/internal/gocmd"
	"github.com/bccfilkom-be/bccli/internal/template"
	"github.com/spf13/cobra"
)

var Framework string

func init() {
	initCmd.Flags().StringVar(&Framework, "framework", "chi", "web framework of choice, [chi]")
	RootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new Go REST server project structure.",
	Long:  "Bootstraps a new Go REST server project by generating the required files\nand directories for running a simple server.",
	Args:  cobra.ExactArgs(1),
	RunE:  _init,
}

func _init(cmd *cobra.Command, args []string) error {
	name := args[0]
	_framework, err := framework.NewFramework(Framework)
	if err != nil {
		return err
	}

	if file.Exist("go.mod") {
		return errors.New("project already initialized")
	}
	if err := gocmd.Init(name); err != nil {
		return err
	}
	if err := framework.Main(_framework); err != nil {
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

	if err := gocmd.Tidy(); err != nil {
		return err
	}

	fmt.Printf("project %s successfully initialized", name)
	return nil
}
