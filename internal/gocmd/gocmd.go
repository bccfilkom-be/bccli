package gocmd

import "os/exec"

// Init go module
func Get(_package string) error {
	return execute("get", "-u", _package)
}

// Init go module
func Init(_package string) error {
	return execute("mod", "init", _package)
}

// Print go version
func Version() (string, error) {
	return output("version")
}

// Execute go command and return error if any
func execute(args ...string) error {
	return exec.Command("go", args...).Run()
}

// Same as execute, but also return stdout
func output(args ...string) (string, error) {
	out, err := exec.Command("go", args...).Output()
	return string(out), err
}
