package main

import (
	"github.com/bccfilkom-be/bccli/cmd"
	_ "github.com/bccfilkom-be/bccli/cmd/domain"
	_ "github.com/bccfilkom-be/bccli/cmd/infra"
)

func main() {
	cmd.Execute()
}
