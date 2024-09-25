package gocmd

import "os/exec"

// Init go module
func Get(pkgs ...string) error {
	for _, pkg := range pkgs {
		err := execute("get", "-u", pkg)
		if err != nil {
			return err
		}
	}
	return nil
}

// Init go module
func Init(pkg string) error {
	return execute("mod", "init", pkg)
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
