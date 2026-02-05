package main

import (
	"os"

	"github.com/1joat/gob/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
