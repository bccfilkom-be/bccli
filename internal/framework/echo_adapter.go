package framework

import (
	"github.com/bccfilkom-be/bccli/internal/file"
	"github.com/bccfilkom-be/bccli/internal/template"
)

type echo struct{}

func (f *echo) Main() error {
	_file, err := file.Create("cmd/main.go")
	if err != nil {
		return err
	}
	if err := template.Execute(_file, "echo_main", nil); err != nil {
		return err
	}
	return nil
}

func (f *echo) Interface() error {
	return nil
}
