/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/template"
	"be-cli/util"
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
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

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

		file, err := util.CreateFile("cmd/main.go")
		if err != nil {
			return err
		}
		fmt.Printf(successCreateFile, "cmd/main.go")

		//give content to main.go
		fileString, err := template.GetFileString("file-template/main.tmpl")
		if err != nil {
			return err
		}

		err = util.ExecuteTemplate(nil, "main.tmpl", fileString, file)
		if err != nil {
			return err
		}

		if err := os.MkdirAll("domain", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "domain")

		if err := os.MkdirAll("infrastructure", os.ModePerm); err != nil {
			return err
		} else {
			fmt.Printf(successCreateDirectory, "infrastructure")
		}

		if err := os.MkdirAll("middleware", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "middleware")

		file, err = util.CreateFile("rest/rest.go")
		if err != nil {
			return err
		}
		fmt.Printf(successCreateFile, "rest/rest.go")

		//give content to rest.go
		_, err = file.Write([]byte("package rest"))
		if err != nil {
			return err
		}

		if err := os.MkdirAll("util", os.ModePerm); err != nil {
			return err
		}
		fmt.Println("successed: Make directory util")

		if _, err := util.CreateFile(".env"); err != nil {
			return err
		}
		fmt.Printf(successCreateFile, ".env")

		file, err = util.CreateFile(".gitignore")
		if err != nil {
			return err
		}

		_, err = file.Write([]byte(".env"))
		if err != nil {
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

		//create mark file and give content
		file, err = util.CreateFile(".mark")
		if err != nil {
			return err
		}

		_, err = file.Write([]byte("This project was created by be-cli. Don't delete this file if you still want to use the be-cli feature"))
		if err != nil {
			return err
		}

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
