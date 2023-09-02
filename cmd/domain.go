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
	Use:   "domain",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
