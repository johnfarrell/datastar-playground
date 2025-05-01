package main

import (
	"github.com/johnfarrell/datastar-playground/internal/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
