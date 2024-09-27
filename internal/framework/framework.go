package framework

import (
	"fmt"

	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/gocmd"
	"github.com/bccfilkom-be/bccli/internal/template"
)

type Framework int

const (
	NET Framework = iota
	CHI
	MUX
	FIBER
	GIN
)

func (f Framework) String() string {
	switch f {
	case NET:
		return "net"
	case CHI:
		return "chi"
	case MUX:
		return "mux"
	case FIBER:
		return "fiber"
	case GIN:
		return "gin"
	default:
		return "unknown"
	}
}

func (f Framework) Package() string {
	switch f {
	case CHI:
		return "github.com/go-chi/chi/v5"
	case MUX:
		return "github.com/gorilla/mux"
	case FIBER:
		return "github.com/gofiber/fiber/v2"
	case GIN:
		return "github.com/gin-gonic/gin"
	default:
		return ""
	}
}

func NewFramework(f string) (Framework, error) {
	switch f {
	case "net":
		return NET, nil
	case "chi":
		return CHI, nil
	case "mux":
		return MUX, nil
	case "fiber":
		return FIBER, nil
	case "gin":
		return GIN, nil
	default:
		return -1, ErrNotExist
	}
}

func Main(f Framework) error {
	_file, err := file.Create("cmd/api/main.go")
	if err != nil {
		return err
	}
	if err := gocmd.Get(f.Package()); err != nil {
		return err
	}
	if err := template.Execute(_file, fmt.Sprintf("%s_main", f), nil); err != nil {
		return err
	}
	return nil
}
