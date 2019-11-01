package main

import (
	"os"

	"github.com/miha-pavel/db_size/cmd"

	// Make sure all your modules and metricsets are linked in this file
	_ "github.com/miha-pavel/db_size/include"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
