/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/template"
	"be-cli/util"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	database string
)

type Database struct {
	Type       string
	DataSource string
	Package    string
	Command    string
}

// infraCmd represents the infra command
var infraCmd = &cobra.Command{
	Use:   "infra",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error{
		if database != "" {
			var data Database
			if database == "mysql" {
				data = Database{
					Type:       database,
					DataSource: "\"%s:%s@tcp(%s)/%s\",os.Getenv(\"DB_USER\"),os.Getenv(\"DB_PASS\"),os.Getenv(\"DB_HOST\"),os.Getenv(\"DB_NAME\"),",
					Package:    "github.com/go-sql-driver/mysql",
					Command:   	"go get github.com/go-sql-driver/mysql",
				}
			} else if database == "postgresql" {
				data = Database{
					Type:       "postgres",
					DataSource: "\"postgresql://%s:%s@%s/%s?sslmode=disable\",os.Getenv(\"DB_USER\"),os.Getenv(\"DB_PASS\"),os.Getenv(\"DB_HOST\"),os.Getenv(\"DB_NAME\"),",
					Package:    "github.com/lib/pq",
					Command:   	"go get github.com/lib/pq",
				}
			} else {
				return errors.New("Specified your database type. ex: mysql,postgresql")
			}

			_, err := util.ExecuteCommand(data.Command)
			if err != nil {
				return err
			} else {
				fmt.Println("successed: Installing "+data.Type+" driver")
			}

			file, err := util.CreateFile("infrastructure/" + database + ".go")
			if err != nil {
				return err
			} else {
				fmt.Println("successed: Make file infrastructure/" + database + ".go")
			}

			fileString, err := template.GetFileString("file-template/sql.tmpl")
			if err != nil {
				return err
			}

			err = util.ExecuteTemplate(data, "sql.tmpl", fileString, file)
			if err != nil {
				return err
			} else {
				fmt.Println("successed: Create " + database + " database connection")
			}
		}else{
			cmd.Help()
		}
		return nil
	},
}

func init() {
	infraCmd.Flags().StringVarP(&database, "db", "d", "", "Flag to generate database connection")

	generateCmd.AddCommand(infraCmd)
}
