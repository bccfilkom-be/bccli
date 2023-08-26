/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/util"
	"errors"
	"fmt"
	"os"

	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

const (
	successCreateFile      string = "successed: Make file %s\n"
	successCreateDirectory string = "successed: Make Directory %s\n"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "generate go project for building REST API",
	Long: `This command is used for make project with name that you specified.
This project purposed for building REST API with clean architecture inspired by Uncle Bob`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("specified your project name")
		}

		str := stringy.New(args[0])

		//make new project directory
		projectName := str.SnakeCase().ToLower()
		if err := os.MkdirAll(projectName, os.ModePerm); err != nil {
			return err
		}

		//run go mod init
		if err := os.Chdir(projectName); err != nil {
			return err
		}

		command := "go mod init " + projectName
		_, err := util.ExecuteCommand(command)
		if err != nil {
			return err
		}
		fmt.Println("successed: Make go project")

		//gin init
		fmt.Println("installing gin ...")

		_, err = util.ExecuteCommand("go get -u github.com/gin-gonic/gin")
		if err != nil {
			return err
		}
		fmt.Println("successed: Installing gin")

		//bootstrapping
		if err := os.MkdirAll("app", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "app")

		if _, err := util.CreateFile("cmd/main.go"); err != nil {
			return err
		}
		fmt.Printf(successCreateFile, "cmd/main.go")

		if err := os.MkdirAll("domain", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "domain")

		if err := os.MkdirAll("deploy", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "deploy")

		if err := os.MkdirAll("infrastructure", os.ModePerm); err != nil {
			return err
		} else {
			fmt.Printf(successCreateDirectory, "infrastructure")
		}

		if err := os.MkdirAll("middleware", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "middleware")

		if _, err := util.CreateFile("rest/rest.go"); err != nil {
			return err
		}
		fmt.Printf(successCreateFile, "rest/rest.go")

		if err := os.MkdirAll("util", os.ModePerm); err != nil {
			return err
		}
		fmt.Println("successed: Make directory util")

		if _, err := util.CreateFile(".env"); err != nil {
			return err
		}
		fmt.Printf(successCreateFile, ".env")

		if _, err := util.CreateFile(".gitignore"); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(successCreateFile, ".gitignore")

		if _, err := util.CreateFile("README.md"); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(successCreateFile, "README.md")

		if _, err := util.CreateFile("Dockerfile"); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(successCreateFile, "Dockerfile")

		if _, err := util.CreateFile("docker-compose.yaml"); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(successCreateFile, "docker-compose.yaml")

		if err := os.Chdir("../"); err != nil {
			return err
		}

		return nil
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
