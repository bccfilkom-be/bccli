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
	Command    string
}

// infraCmd represents the infra command
var infraCmd = &cobra.Command{
	Use:   "infra [flags]",
	Short: "Generate infra [flags] components",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {

		if cmd.Flags().NFlag() == 0 {
			fmt.Println("error: please specified your flag. ex: --db=mysql")
			return
		}
    
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
				fmt.Println("error: please specified your database type. ex: --db=mysql")
				return
			}

			_, err := util.ExecuteCommand(data.Command)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				fmt.Println("successed: Installing "+data.Type+" driver")
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
		}else{
			cmd.Help()
		}
	},
}

func init() {
	infraCmd.Flags().StringVarP(&database, "db", "d", "", "Flag to generate database connection. infra-spesific-type: mysql,postgresql")
  
	rootCmd.AddCommand(infraCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// infraCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// infraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
