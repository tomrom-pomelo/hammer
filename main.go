package main

import (
	"os"

	"github.com/tomrom-pomelo/hammer/cmd"
)

var version string

func main() {
	cmd.Version = version

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
