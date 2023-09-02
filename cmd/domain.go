/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/template"
	"be-cli/util"
	"errors"
	"fmt"

	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

type Data struct {
	Domain string
}

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain <domain-name>",
	Short: "Generate a new domain or entity for your REST API project.",
	Long: `This command allows you to create a new domain or entity within your REST API project. A domain entity represents a specific resource or data structure in your API, such as User, Product, or Order. Customize the generated code according to your project's specific requirements.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return errors.New("Specified your domain name")
		}

		str := stringy.New(args[0])
		domainName := str.SnakeCase().ToLower()
		file, err := util.CreateFile("domain/" + domainName + ".go")
		if err != nil {
			return err
		} else {
			fmt.Println("successed: Make file domain/" + domainName + ".go")
		}

		data := Data{
			Domain: str.CamelCase(),
		}

		fileString, err := template.GetFileString("file-template/domain.tmpl")
		if err != nil {
			return err
		}

		err = util.ExecuteTemplate(data, "domain.tmpl", fileString, file)
		if err != nil {
			return err
		} else {
			fmt.Println("successed: Create " + domainName + " domain")
		}
		return nil
	},
}

func init() {
	generateCmd.AddCommand(domainCmd)
}
