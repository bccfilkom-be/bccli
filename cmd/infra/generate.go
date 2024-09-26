package infra

import (
	"fmt"
	"strings"

	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/gocmd"
	"github.com/bccfilkom-be/bccli/internal/template"
	"github.com/spf13/cobra"
)

var genCmd = &cobra.Command{
	Use:   "generate <service>",
	Short: "Generate app 3rd party service client.",
	Long:  "Generate 3rd party service client to be used. Available values are mysql, mariadb and postgresql",
	Args:  cobra.ExactArgs(1),
	RunE:  gen,
}

func init() {
	infraCmd.AddCommand(genCmd)
}

var (
	services = map[string][]string{
		"mysql":      {"github.com/go-sql-driver/mysql", "github.com/jmoiron/sqlx"},
		"mariadb":    {"github.com/go-sql-driver/mysql", "github.com/jmoiron/sqlx"},
		"postgresql": {"github.com/jackc/pgx/v5"},
	}
)

func gen(cmd *cobra.Command, args []string) error {
	service := strings.ToLower(args[0])
	pkg, ok := services[service]
	if !ok {
		return ErrNotExist
	}
	if err := gocmd.Get(pkg...); err != nil {
		return err
	}
	file, err := file.Create(fmt.Sprintf("internal/infra/%s.go", service))
	if err != nil {
		return err
	}
	if err := template.Execute(file, service, nil); err != nil {
		return err
	}

	fmt.Printf("service %s succesfully generated", service)
	return nil
}
