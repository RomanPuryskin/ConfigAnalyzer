package main

import (
	"os"

	"github.com/configAnalyzer/cmd/cli"
	_ "github.com/spf13/cobra"
)

func main() {

	if err := cli.RootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}
