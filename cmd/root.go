package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// MakeRootCmd creates a new instance of the Moducate root command. All commands
// (serve, migrate, help, etc.) are children of this command.
func MakeRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "moducate",
		Short: "Run and manage Moducate",
	}
}

var rootCmd = MakeRootCmd()

// Execute performs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		if _, err := fmt.Fprintln(os.Stderr, err); err != nil {
			panic(err.Error())
		}
		os.Exit(1)
	}
}
