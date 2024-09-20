package template

import (
	"embed"
	"os"
	_template "text/template"
)

//go:embed tmpl
var fs embed.FS

func Execute(file *os.File, filename string, data interface{}) error {
	filename += ".tmpl"
	_content, err := content(filename)
	if err != nil {
		return err
	}
	tmpl, err := _template.New(filename).Parse(_content)
	if err != nil {
		return err
	}
	if err := tmpl.Execute(file, data); err != nil {
		return err
	}
	return nil
}

func content(file string) (string, error) {
	fileString, err := fs.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(fileString), nil
}
