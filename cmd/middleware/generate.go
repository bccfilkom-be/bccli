package middleware

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/template"
	"github.com/gobeam/stringy"
	"github.com/spf13/cobra"
)

type Data struct {
	Middleware string
}

type Framework string

const (
	FrameworkGin   Framework = "gin"
	FrameworkFiber Framework = "fiber"
	FrameworkChi   Framework = "chi"
)

var moduleByFramework = map[Framework][]string{
	FrameworkGin:   {"github.com/gin-gonic/gin"},
	FrameworkFiber: {"github.com/gofiber/fiber/v2"},
	FrameworkChi:   {"github.com/go-chi/chi/v5"},
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

	middlewareFile, err := file.Create("internal/middleware/" + middlewareName + ".go")
	if err != nil {
		return err
	}

	fw, err := detectFramework(".")
	if err != nil {
		return err
	}

	data := Data {
		Middleware: middlewareName,
	}

	switch fw {
	case FrameworkChi:
		err = template.Execute(middlewareFile, "ChiMiddleware", data)
	case FrameworkGin:
		err = template.Execute(middlewareFile, "GinMiddleware", data)
	case FrameworkFiber:
		err = template.Execute(middlewareFile, "FiberMiddleware", data)
	}

	if err != nil {
		return err
	}

	fmt.Printf("middleware %s succesfully generated", middlewareName)

	return nil
}

func detectFramework(projectDir string) (Framework, error) {
	goModPath := filepath.Join(projectDir, "go.mod")
	b, err := os.ReadFile(goModPath)
	if err != nil {
		return "", fmt.Errorf("read go.mod: %w", err)
	}
	text := string(b)

	seen := map[Framework]struct{}{}
	for fw, mods := range moduleByFramework {
		for _, m := range mods {
			if strings.Contains(text, m) {
				seen[fw] = struct{}{}
				break
			}
		}
	}

	switch len(seen) {
	case 0:
		return "", errors.New("no supported framework modules found in go.mod")
	case 1:
		for fw := range seen {
			return fw, nil
		}
	default:
		var names []string
		for fw := range seen {
			names = append(names, string(fw))
		}
		sort.Strings(names)
		return "", fmt.Errorf("multiple frameworks detected in go.mod: %s. use --framework to override", strings.Join(names, ", "))
	}
	return "", errors.New("unreachable")
}
