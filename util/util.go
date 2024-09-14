package util

import (
	"errors"
	"text/template"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CreateFile(name string) (*os.File, error) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		if strings.Contains(name, "/") {
			path := name[:strings.LastIndex(name, "/")]

			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return nil, err
			}
		}

		file, err := os.Create(name)
		if err != nil {
			return nil, err
		}
		return file, nil
	} else {
		return nil, errors.New("file already exist")
	}
}

func ExecuteCommand(command string) (string, error) {
	var output []byte
	var err error

	if runtime.GOOS == "windows" {
		output, err = exec.Command("cmd", "/C", command).Output()
	} else {
		output, err = exec.Command("bash", "-c", command).Output()
	}

	return string(output), err
}

func ExecuteTemplate(data interface{}, templateFile string, fileString string, file *os.File) error {
	tmplHandler, err := template.New(templateFile).Parse(fileString)
	if err != nil {
		return err
	}
	err = tmplHandler.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}
