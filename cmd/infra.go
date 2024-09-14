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
	Long: `This command allows you to generate infrastructure components for your REST API project. You can specify the type of infrastructure you want to generate using the flag and provide the specific type of infrastructure as an argument. This command streamlines the process of creating essential infrastructure components for your REST API, such as databases, caching systems, or other dependencies. It ensures that your project has the necessary infrastructure in place to support its functionality. Customize the generated infrastructure components according to your project's specific requirements, and  use the appropriate flags and arguments to tailor the generation process.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		if cmd.Flags().NFlag() == 0 {
			return fmt.Errorf("please specified your flag. ex: --db=mysql")
		}

		var data Database

		switch database {
		case "mysql":
			data = Database{
				Type:       database,
				DataSource: "\"%s:%s@tcp(%s)/%s\",os.Getenv(\"DB_USER\"),os.Getenv(\"DB_PASS\"),os.Getenv(\"DB_HOST\"),os.Getenv(\"DB_NAME\"),",
				Package:    "github.com/go-sql-driver/mysql",
				Command:    "go get github.com/go-sql-driver/mysql",
			}
		case "postgresql":
			data = Database{
				Type:       "postgres",
				DataSource: "\"postgresql://%s:%s@%s/%s?sslmode=disable\",os.Getenv(\"DB_USER\"),os.Getenv(\"DB_PASS\"),os.Getenv(\"DB_HOST\"),os.Getenv(\"DB_NAME\"),",
				Package:    "github.com/lib/pq",
				Command:    "go get github.com/lib/pq",
			}
		default:
			return fmt.Errorf("database type not found")
		}

		_, err := util.ExecuteCommand(data.Command)
		if err != nil {
			return err
		} else {
			fmt.Println("successed: Installing " + data.Type + " driver")
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

		return nil
	},
}

func init() {
	infraCmd.Flags().StringVarP(&database, "db", "d", "", "Flag to generate database connection. infra-spesific-type: mysql,postgresql")

	rootCmd.AddCommand(infraCmd)
}
