/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"be-cli/template"
	"be-cli/util"
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
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			fmt.Println("Specified your domain name")
			return
		}

		str := stringy.New(args[0])
		domainName := str.SnakeCase().ToLower()
		file, err := util.CreateFile("domain/" + domainName + ".go")
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Make file domain/" + domainName + ".go")
		}

		data := Data{
			Domain: str.CamelCase(),
		}

		fileString, err := template.GetFileString("file-template/domain.tmpl")
		if err != nil {
			fmt.Println(err)
			return
		}


		err = util.ExecuteTemplate(data, "domain.tmpl", fileString, file)
		if err != nil {
			fmt.Println(err)
			return
		} else {
			fmt.Println("successed: Create " + domainName + " domain")
		}
	},
}

func init() {
	generateCmd.AddCommand(domainCmd)
}
