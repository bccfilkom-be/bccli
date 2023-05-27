/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"be-cli/util"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/gobeam/stringy"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Specified your project name")
			return
		}

		str := stringy.New(args[0])

		//make new project directory
		projectName := str.SnakeCase().ToLower()
		if err := os.MkdirAll(projectName, os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		//run go mod init
		if err := os.Chdir(projectName); err != nil {
			fmt.Println(err)
			return
		}

		command := "go mod init " + projectName
		_, err := util.ExecuteCommand(command)
		if err != nil {
			fmt.Println(err)
			return
		}

		//gin init
		output, err := util.ExecuteCommand("go get -u github.com/gin-gonic/gin")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(output)

		//bootstrapping
		if err := os.MkdirAll("app", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		if err := util.CreateFile("cmd/main.go"); err != nil {
			fmt.Println(err)
			return
		}

		if err := os.MkdirAll("domain", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		if err := os.MkdirAll("deploy", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		if err := os.MkdirAll("infrastructure", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		if err := os.MkdirAll("middleware", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		if err := util.CreateFile("rest/rest.go"); err != nil {
			fmt.Println(err)
			return
		}

		if err := os.MkdirAll("util", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		}

		if err := util.CreateFile(".env"); err != nil {
			fmt.Println(err)
			return
		}

		if err := util.CreateFile(".gitignore"); err != nil {
			fmt.Println(err)
			return
		}

		if err := util.CreateFile("README.md"); err != nil {
			fmt.Println(err)
			return
		}

		if err := util.CreateFile("Dockerfile"); err != nil {
			fmt.Println(err)
			return
		}

		if err := util.CreateFile("docker-compose.yaml"); err != nil {
			fmt.Println(err)
			return
		}

		if err := os.Chdir("../"); err != nil {
			fmt.Println(err)
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
