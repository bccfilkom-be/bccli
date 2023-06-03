package util

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"text/template"
)

func CreateFile(name string) (*os.File,error) {
	_, err := os.Stat(name)

	if os.IsNotExist(err) {
		if strings.Contains(name, "/") {
			path := name[:strings.LastIndex(name, "/")]

			err := os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return nil,err
			}
		}

		file, err := os.Create(name)
		if err != nil {
			return nil,err
		}
		return file,nil
	} else {
		return nil,errors.New("file already exist")
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

func ExecuteTemplate(data interface{}, templateFile string, templatePath string, file *os.File)(error){
	tmplHandler, err := template.New(templateFile).ParseFiles(templatePath)
	if err != nil {
		return err
	}
	err = tmplHandler.Execute(file, data)
	if err != nil {
		return err
	}
	return nil
}