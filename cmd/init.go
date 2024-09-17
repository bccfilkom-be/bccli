package cmd

import (
	"fmt"
	"os"

	"github.com/bccfilkom-be/bccli/util"
  "github.com/bccfilkom-be/bccli/template"
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

		//BOOTSTRAPING
		//create directory app
		if err := os.MkdirAll("app", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "app")

		//create directory cmd, file cmd/main.go, give content to main.go
		file, err := util.CreateFile("cmd/main.go")
		if err != nil {
			return err
		}
		fmt.Printf(successCreateFile, "cmd/main.go")

		fileString, err := template.GetFileString("file-template/main.tmpl")
		if err != nil {
			return err
		}

		err = util.ExecuteTemplate(nil, "main.tmpl", fileString, file)
		if err != nil {
			return err
		}

		//create directory domain
		if err := os.MkdirAll("domain", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "domain")

		//create directory infra
		if err := os.MkdirAll("infrastructure", os.ModePerm); err != nil {
			return err
		} else {
			fmt.Printf(successCreateDirectory, "infrastructure")
		}

		//create directory middleware
		if err := os.MkdirAll("middleware", os.ModePerm); err != nil {
			return err
		}
		fmt.Printf(successCreateDirectory, "middleware")

		//create directory rest, file rest/rest.go, give content to rest.go
		file, err = util.CreateFile("middleware/rest/rest.go")
		if err != nil {
			return err
		}
		fmt.Printf(successCreateFile, "rest/rest.go")

		_, err = file.Write([]byte("package rest"))
		if err != nil {
			return err
		}

		//create directory util
		if err := os.MkdirAll("util", os.ModePerm); err != nil {
			return err
		}
		fmt.Println("successed: Make directory util")

		//create file .env
		if _, err := util.CreateFile(".env"); err != nil {
			return err
		}
		fmt.Printf(successCreateFile, ".env")

		//create file .env.example
		if _, err := util.CreateFile(".env.example"); err != nil {
			return err
		}
		fmt.Printf(successCreateFile, ".env.example")

		//create file .gitignore and give it content
		file, err = util.CreateFile(".gitignore")
		if err != nil {
			return err
		}
		fmt.Printf(successCreateFile, ".gitignore")

		_, err = file.Write([]byte(".env"))
		if err != nil {
			return err
		}

		//create file README.md
		if _, err := util.CreateFile("README.md"); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(successCreateFile, "README.md")

		//create file Dockerfile
		if _, err := util.CreateFile("Dockerfile"); err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf(successCreateFile, "Dockerfile")

		//create file //docker-compose.yaml
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
}
