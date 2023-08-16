/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/template"
	"be-cli/util"
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
	Run: func(cmd *cobra.Command, args []string) {
		if database != "" {
			var data Database
			if database == "mysql" {
				_, err := util.ExecuteCommand("go get github.com/go-sql-driver/mysql")
				if err != nil {
					fmt.Println(err)
					return
				} else {
					fmt.Println("successed: Installing mysql driver")
				}
				data = Database{
					Type:       database,
					DataSource: "\"%s:%s@tcp(%s)/%s\",os.Getenv(\"DB_USER\"),os.Getenv(\"DB_PASS\"),os.Getenv(\"DB_HOST\"),os.Getenv(\"DB_NAME\"),",
					Package:    "github.com/go-sql-driver/mysql",
				}
			} else if database == "postgresql" {
				_, err := util.ExecuteCommand("go get github.com/lib/pq")
				if err != nil {
					fmt.Println(err)
					return
				} else {
					fmt.Println("successed: Installing postgresql driver")
				}
				data = Database{
					Type:       database,
					DataSource: "\"postgresql://%s:%s@%s/%s?sslmode=disable\",os.Getenv(\"DB_USER\"),os.Getenv(\"DB_PASS\"),os.Getenv(\"DB_HOST\"),os.Getenv(\"DB_NAME\"),",
					Package:    "github.com/lib/pq",
				}
			} else {
				fmt.Println("Specified your database type. ex: mysql,postgresql")
				return
			}

			file, err := util.CreateFile("infrastructure/" + database + ".go")
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println("successed: Make file infrastructure/" + database + ".go")
			}

			fileString, err := template.GetFileString("file-template/sql.tmpl")
			if err != nil {
				fmt.Println(err)
				return
			}

			err = util.ExecuteTemplate(data, "sql.tmpl", fileString, file)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println("successed: Create " + database + " database connection")
			}
		}
	},
}

func init() {
	infraCmd.Flags().StringVarP(&database, "db", "d", "", "Flag to generate database connection")

	generateCmd.AddCommand(infraCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
