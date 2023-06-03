package util

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CreateFile(name string) (*os.File, error) {

	if _, err := os.Stat(name); os.IsExist(err) {
		return nil, errors.New("file already exist")
	}

	if strings.ContainsAny(name, "/\\") {
		path := name[:strings.LastIndexAny(name, "/\\")]

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	fl, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return fl, nil
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
