package middleware

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/framework"
	"github.com/bccfilkom-be/bccli/internal/template"
	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

type Data struct {
	Middleware string
}

var moduleByFramework = map[framework.Framework][]string{
	framework.GIN:   {"github.com/gin-gonic/gin"},
	framework.FIBER: {"github.com/gofiber/fiber/v2"},
	framework.CHI:   {"github.com/go-chi/chi/v5"},
}

func init() {
	middlewareCmd.AddCommand(genCmd)
}

var genCmd = &cobra.Command{
	Use:   "generate <middleware>",
	Short: "",
	Long:  "",
	Args:  cobra.ExactArgs(1),
	RunE:  gen,
}

func gen(cmd *cobra.Command, args []string) error {
	str := stringy.New(args[0])
	middlewareName := str.SnakeCase().ToLower()

	fw, err := detectFramework(".")
	if err != nil {
		return err
	}

	middlewareFile, err := file.Create("internal/middleware/" + middlewareName + ".go")
	if err != nil {
		return err
	}

	data := Data{
		Middleware: str.PascalCase().Get(),
	}

	switch fw {
	case framework.CHI:
		err = template.Execute(middlewareFile, "chi_middleware", data)
	case framework.GIN:
		err = template.Execute(middlewareFile, "gin_middleware", data)
	case framework.FIBER:
		err = template.Execute(middlewareFile, "fiber_middleware", data)
	}

	if err != nil {
		return err
	}

	fmt.Printf("middleware %s succesfully generated", middlewareName)

	return nil
}

func detectFramework(projectDir string) (framework.Framework, error) {
	goModPath := filepath.Join(projectDir, "go.mod")
	b, err := os.ReadFile(goModPath)
	if err != nil {
		return -1, err
	}
	text := string(b)

	for fw, mods := range moduleByFramework {
		for _, m := range mods {
			if strings.Contains(text, m) {
				return fw, nil
			}
		}
	}

	return -1, ErrFWNotFound
}
