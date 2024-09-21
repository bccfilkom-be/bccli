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

func NewFramework(f string) Framework {
	switch f {
	case "net":
		return NET
	case "chi":
		return CHI
	case "mux":
		return MUX
	case "fiber":
		return FIBER
	case "gin":
		return GIN
	default:
		return -1
	}
}

func Main(f Framework) error {
	_file, err := file.Create("cmd/api/main.go")
	if err != nil {
		return err
	}
	if err := gocmd.Get(chi); err != nil {
		return err
	}
	if err := template.Execute(_file, fmt.Sprintf("%s_main", f), nil); err != nil {
		return err
	}
	return nil
}
