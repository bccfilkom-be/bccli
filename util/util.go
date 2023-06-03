package util

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CreateFile(name string) (file *os.File, err error) {
	_, err = os.Stat(name)

	if os.IsNotExist(err) {
		if strings.ContainsAny(name, "/\\") {
			path := name[:strings.LastIndexAny(name, "/\\")]

			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return nil, err
			}
		}

		file, err = os.Create(name)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("file already exist")
	}

	return file, nil
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
