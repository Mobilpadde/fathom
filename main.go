package main

import (
	"fmt"
	"os"

	"analytome/pkg/cli"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	err := cli.Run(version, commit, date)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	os.Exit(0)
}
