package main

import (
	"github.com/bccfilkom-be/bccli/cmd"
	_ "github.com/bccfilkom-be/bccli/cmd/domain"
	_ "github.com/bccfilkom-be/bccli/cmd/infra"
	_ "github.com/bccfilkom-be/bccli/cmd/middleware"
)

func main() {
	cmd.Execute()
}
