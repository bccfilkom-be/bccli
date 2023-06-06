/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/util"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

type appOptions struct {
	handler    bool
	repository bool
	usecase    bool
}

var appOpts appOptions

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:          "app [<name>]",
	Short:        "Generate app [<name>] components",
	Long:         ``,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		// Create and Change current working directory to the new app
		appName := stringy.New(args[0]).SnakeCase().ToLower()
		if err := os.MkdirAll("app/"+appName, os.ModePerm); err != nil {
			return err
		}

		if err := os.Chdir("app/" + appName); err != nil {
			return err
		}

		// Create and write all the required folders and .go files for the app

		// Generate all components if no flag are specified
		if !appOpts.handler && !appOpts.repository && !appOpts.usecase {
			if err := generateComponent(appName, "handler"); err != nil {
				return err
			}
			if err := generateComponentWithMock(appName, "repository"); err != nil {
				return err
			}
			if err := generateComponentWithMock(appName, "usecase"); err != nil {
				return err
			}
			return nil
		}

		// Generate component based on specified flag
		if appOpts.handler {
			if err := generateComponent(appName, "handler"); err != nil {
				return err
			}
		}
		if appOpts.repository {
			if err := generateComponentWithMock(appName, "repository"); err != nil {
				return err
			}
		}
		if appOpts.usecase {
			if err := generateComponentWithMock(appName, "usecase"); err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	appCmd.Flags().BoolVar(&appOpts.handler, "handler", false, "Generate handler component")
	appCmd.Flags().BoolVar(&appOpts.repository, "repository", false, "Generate repository component")
	appCmd.Flags().BoolVar(&appOpts.usecase, "usecase", false, "Generate usecase component")

	generateCmd.AddCommand(appCmd)
}

// Generate a Component along with its folder. This also include file and test file
func generateComponent(appName string, componentName string) error {
	pathStr := path.Join(componentName, fmt.Sprintf("%s_%s.go", appName, componentName))
	if err := createAndWriteFile(pathStr, componentName); err != nil {
		return err
	}

	pathStr = pathStr[:strings.Index(pathStr, ".go")] + "_test.go"
	if err := createAndWriteFile(pathStr, componentName); err != nil {
		return err
	}

	return nil
}

// Same as generateComponent() but this include mock file
func generateComponentWithMock(appName string, componentName string) error {
	if err := generateComponent(appName, componentName); err != nil {
		return err
	}

	pathStr := path.Join(componentName, fmt.Sprintf("%s_%s_mock.go", appName, componentName))
	if err := createAndWriteFile(pathStr, componentName); err != nil {
		return err
	}

	return nil
}

// Create and Write File components that include package name inside
func createAndWriteFile(filePath string, componentName string) error {
	fi, err := util.CreateFile(filePath)
	if err != nil {
		return err
	}

	_, err = fi.Write([]byte("package " + componentName))
	if err != nil {
		return err
	}

	return nil
}
