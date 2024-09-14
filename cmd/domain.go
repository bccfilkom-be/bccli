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
	Use:   "domain [<name>]",
	Short: "Generate domain [<name>] components",
	Long: ``,
	RunE: func(cmd *cobra.Command, args []string) error{
		if len(args) == 0 {
			return errors.New("specified your domain name")
		}

		str := stringy.New(args[0])
		domainName := str.CamelCase()
		file, err := util.CreateFile("domain/" + domainName + ".go")
		if err != nil {
			return err
		} else {
			fmt.Println("successed: Make file domain/" + domainName + ".go")
		}

		data := Data{
			Domain: domainName,
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
	rootCmd.AddCommand(domainCmd)
}
