package domain

import (
	"errors"
	"fmt"
	"os"
	"path"

	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/template"
	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

type Data struct {
	Domain       string
	DomainStruct string
	DBDriver     string
	Database     string
	Module       string
}

type appOptions struct {
	handler    bool
	usecase    bool
	repository string
}

var (
	appOpts  appOptions
	database string
	dbMap    = map[string]struct {
		name   string
		module string
	}{
		"mysql": {
			name:   "MySQL",
			module: `"github.com/jmoiron/sqlx"`,
		},
		"mariadb": {
			name:   "MariaDB",
			module: `"github.com/jmoiron/sqlx"`,
		},
		"postgresql": {
			name:   "PostgreSQL",
			module: `"github.com/jackc/pgx/v5"`,
		},
	}
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate domain components like handler, usecase, and repository.",
	Long: `Generate domain components like handler, usecase, and repository.
By default it generate all component at once, but you can choose one or multiple.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("specified your domain name")
		}

		str := stringy.New(args[0])
		domainName := str.SnakeCase().ToLower()

		if appOpts.repository != "" {
			database = appOpts.repository
			err := generateComponent(domainName, "repository", database)
			if err != nil {
				return err
			}

			fmt.Printf("domain %s succesfully generated", domainName)
			return nil
		}

		if appOpts.handler {
			err := generateComponent(domainName, "handler", "")
			if err != nil {
				return err
			}

			fmt.Printf("domain %s succesfully generated", domainName)
			return nil
		}

		if appOpts.usecase {
			err := generateComponent(domainName, "usecase", "")
			if err != nil {
				return err
			}

			fmt.Printf("domain %s succesfully generated", domainName)
			return nil
		}

		err := generateAllComponent(domainName, database)
		if err != nil {
			return err
		}

		file, err := file.Create("internal/domain/" + domainName + ".go")
		if err != nil {
			return err
		}

		data := Data{
			Domain: str.CamelCase().Get(),
		}

		err = template.Execute(file, "domain", data)
		if err != nil {
			return err
		}

		if err := os.MkdirAll("internal/"+domainName, os.ModePerm); err != nil {
			return err
		}

		fmt.Printf("domain %s succesfully generated", domainName)

		return nil
	},
}

func init() {
	generateCmd.Flags().BoolVarP(&appOpts.handler, "handler", "H", false, "Generate handler")
	generateCmd.Flags().BoolVarP(&appOpts.usecase, "usecase", "U", false, "Generate usecase")
	generateCmd.Flags().StringVarP(&appOpts.repository, "repository", "R", "", "Generate repository. choose specific database, ex: mysql,postgresql")
	generateCmd.Flags().StringVarP(&database, "database", "d", "", "Flag to generate repository file. choose specific database, ex: mysql,postgresql")

	domainCmd.AddCommand(generateCmd)
}

func generateAllComponent(domainName, database string) error {
	components := []string{"repository", "handler", "usecase"}
	for _, component := range components {
		if err := generateComponent(domainName, component, database); err != nil {
			return err
		}
	}

	return nil
}

func generateComponent(domainName, componentName, database string) error {
	pathMap := map[string]string{
		"repository": "repository",
		"handler":    "interface/rest",
		"usecase":    "usecase",
	}

	dirPath := path.Join("internal/", domainName, pathMap[componentName])
	if componentName == "repository" && database == "" {
		err := os.MkdirAll(dirPath, os.ModePerm)
		return err
	}

	if componentName == "repository" && database != "" {
		_, err := os.Stat(fmt.Sprintf("internal/infra/%s.go", database))
		if os.IsExist(err) {
			return errors.New("file already exist")
		}

		if err == nil {
			filePath := path.Join(dirPath, fmt.Sprintf("%s.go", database))
			componentName = "repoWithDb"

			return createAndWriteFile(filePath, componentName, domainName)
		}

		return errors.New("database not found")
	}

	filePath := path.Join(dirPath, fmt.Sprintf("%s.go", domainName))
	return createAndWriteFile(filePath, componentName, domainName)
}

func createAndWriteFile(filePath, componentName, domainName string) error {
	str := stringy.New(domainName)

	dbMap := dbMap[database]

	data := Data{
		Domain:       str.CamelCase().Get(),
		DomainStruct: str.SnakeCase().ToLower(),
		Database:     dbMap.name,
	}

	if componentName == "repoWithDb" {
		data.Module = dbMap.module

		switch database {
		case "mysql", "mariadb":
			data.DBDriver = "sqlx.DB"
		case "postgresql":
			data.DBDriver = "pgx.Conn"
		default:
			return fmt.Errorf("database %s not found", database)
		}
	}

	file, err := file.Create(filePath)
	if err != nil {
		return err
	}

	err = template.Execute(file, componentName, data)
	if err != nil {
		return err
	}

	return nil
}
