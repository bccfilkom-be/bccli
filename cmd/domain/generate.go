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
}

type appOptions struct {
	handler    bool
	repository string
	usecase    bool
}

var (
	appOpts  appOptions
	database string
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
		file, err := file.Create("internal/domain/" + domainName + ".go")
		if err != nil {
			return err
		}

		data := Data{
			Domain: str.CamelCase(),
		}

		err = template.Execute(file, "domain", data)
		if err != nil {
			return err
		}

		if err := os.MkdirAll("internal/"+domainName, os.ModePerm); err != nil {
			return err
		}

		if appOpts.repository != "" {
			err := generateComponent(domainName, "repository", appOpts.repository)
			return err
		}

		if appOpts.handler {
			err := generateComponent(domainName, "handler", "")
			return err
		}

		if appOpts.usecase {
			err := generateComponent(domainName, "usecase", "")
			return err
		}

		fmt.Println("database: ", database)

		fmt.Println("Generate all component")
		err = generateAllComponent(domainName, database)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	generateCmd.Flags().BoolVarP(&appOpts.handler, "handler", "H", false, "Generate handler")
	generateCmd.Flags().BoolVarP(&appOpts.usecase, "usecase", "U", false, "Generate usecase")
	generateCmd.Flags().StringVarP(&appOpts.repository, "repository", "R", "", "Generate repository. choose specific database, ex: mysql,postgresql")
	generateCmd.Flags().StringVarP(&database, "database", "d", "", "Flag to generate repository file. choose specific database, ex: mysql,postgresql")

	DomainCmd.AddCommand(generateCmd)
}

func generateAllComponent(domainName, database string) error {
	components := []string{"handler", "repository", "usecase"}
	for _, component := range components {
		if err := generateComponent(domainName, component, database); err != nil {
			return err
		}
	}

	return nil
}

func generateComponent(domainName, componentName, database string) error {
	pathMap := map[string]string{
		"handler":    "interface/rest",
		"repository": "repository",
		"usecase":    "usecase",
	}

	dirPath := path.Join("internal/", domainName, pathMap[componentName])
	if componentName == "repository" && database == "" {
		err := os.MkdirAll(dirPath, os.ModePerm)
		return err
	}

	if componentName == "repository" && database != "" {
		filePath := path.Join(dirPath, fmt.Sprintf("%s.go", database))
		return createAndWriteFile(filePath, componentName, domainName)
	}

	filePath := path.Join(dirPath, fmt.Sprintf("%s.go", domainName))
	return createAndWriteFile(filePath, componentName, domainName)
}

func createAndWriteFile(filePath, componentName, domainName string) error {
	file, err := file.Create(filePath)
	if err != nil {
		return err
	}

	str := stringy.New(domainName)

	data := Data{
		Domain:       str.CamelCase(),
		DomainStruct: str.SnakeCase().ToLower(),
	}

	err = template.Execute(file, componentName, data)
	if err != nil {
		return err
	}

	return nil
}
