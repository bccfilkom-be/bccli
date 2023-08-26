/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/util"
	"fmt"
	"os"

	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "generate go project for building REST API",
	Long:  `This command is used for make project with name that you specified.
This project purposed for building REST API with clean architecture inspired by Uncle Bob`,
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
		} else {
			fmt.Println("successed: Make go project")
		}

		//gin init
		_, err = util.ExecuteCommand("go get -u github.com/gin-gonic/gin")
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Installing gin")
		}

		//bootstrapping
		if err := os.MkdirAll("app", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make directory app")
		}

		if _, err := util.CreateFile("cmd/main.go"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file cmd/main.go")
		}

		if err := os.MkdirAll("domain", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make directory domain")
		}

		if err := os.MkdirAll("deploy", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make directory deploy")
		}

		if err := os.MkdirAll("infrastructure", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make directory infrastructure")
		}

		if err := os.MkdirAll("middleware", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make directory middleware")
		}

		if _, err := util.CreateFile("rest/rest.go"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file rest/rest.go")
		}

		if err := os.MkdirAll("util", os.ModePerm); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make directory util")
		}

		if _, err := util.CreateFile(".env"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file .env")
		}

		if _, err := util.CreateFile(".gitignore"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file .gitignore")
		}

		if _, err := util.CreateFile("README.md"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file README.md")
		}

		if _, err := util.CreateFile("Dockerfile"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file Dockerfile")
		}

		if _, err := util.CreateFile("docker-compose.yaml"); err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file docker-compose.yaml")
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
