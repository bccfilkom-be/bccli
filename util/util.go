package util

import (
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func CreateFile(name string) (*os.File, error) {
	if strings.ContainsAny(name, "/\\") {
		path := name[:strings.LastIndexAny(name, "/\\")]

		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return nil, err
		}
	}

	fl, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
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
