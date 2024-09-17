package shell

import (
	"os/exec"
	"runtime"
)

func Execute(command string) (string, error) {
	var output []byte
	var err error
	switch runtime.GOOS {
	case "windows":
		output, err = exec.Command("cmd", "/C", command).Output()
	default:
		output, err = exec.Command("bash", "-c", command).Output()
	}
	return string(output), err
}
