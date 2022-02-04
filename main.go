package main

import (
	"os"

	"github.com/moducate/moducate/cmd"
)

func main() {
	cmd.Execute(cmd.RootCmd, os.Stderr)
}
