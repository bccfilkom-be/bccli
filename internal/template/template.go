package template

import (
	"embed"
	"fmt"
	"os"
	_template "text/template"
)

//go:embed tmpl
var fs embed.FS

func Execute(file *os.File, filename string, data interface{}) error {
	filename = fmt.Sprint(filename, ".tmpl")
	tmpl, err := _template.New(filename).ParseFS(fs, fmt.Sprintf("tmpl/%s", filename))
	if err != nil {
		return err
	}
	if err := tmpl.Execute(file, data); err != nil {
		return err
	}
	return nil
}
