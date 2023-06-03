/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/util"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// appCmd represents the app command
var appCmd = &cobra.Command{
	Use:          "app [<name>]",
	Short:        "Generate app [<name>] components",
	Long:         ``,
	Args:         cobra.ExactArgs(1),
	SilenceUsage: true,

	RunE: func(cmd *cobra.Command, args []string) error {
		appName := args[0]

		if err := os.Chdir("app"); err != nil {
			return errors.New("unable to locate app folder in the current working directory")
		}

		// Create and write all the required folders and .go files for the app
		dirNames := [3]string{"handler", "repository", "usecase"}

		for _, dirName := range dirNames {
			dirNamePath := filepath.Join(appName, dirName, fmt.Sprintf("%s_%s.go", appName, dirName))
			if err := createAndWriteFile(dirNamePath, dirName); err != nil {
				return err
			}

			dirNamePath = filepath.Join(appName, dirName, fmt.Sprintf("%s_%s_test.go", appName, dirName))
			if err := createAndWriteFile(dirNamePath, dirName); err != nil {
				return err
			}
		}

		for _, dirName := range dirNames[1:] {
			dirNamePath := filepath.Join(appName, dirName, fmt.Sprintf("%s_%s_mock.go", appName, dirName))
			if err := createAndWriteFile(dirNamePath, dirName); err != nil {
				return err
			}
		}

		return nil
	},
}

func init() {
	// appCmd.Flags().StringVarP(&name, "name", "n", "", "sets the app's generated name")
	// if err := appCmd.MarkFlagRequired("name"); err != nil {
	// 	fmt.Println(err.Error())
	// }

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.AddCommand(appCmd)
}

func createAndWriteFile(filePath, packageName string) error {
	fi, err := util.CreateFile(filePath)
	if err != nil {
		return err
	}

	_, err = fi.Write([]byte("package " + packageName))
	if err != nil {
		return err
	}

	return nil
}
