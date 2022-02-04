package cmd

import (
	"github.com/moducate/moducate/internal/db"
	"github.com/spf13/cobra"
)

// MakeMigrateCmd creates a new instance of the `moducate migrate` command. This
// command performs PostgreSQL database migrations.
func MakeMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "migrate <DSN>",
		Short: "Performs Moducate's PostgreSQL database migrations",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				cmd.PrintErrln("Please provide exactly 1 argument for the DSN (PostgreSQL connection string)!")
				return
			}

			applied, err := db.Migrate(args[0])

			if err != nil {
				cmd.PrintErrln("Failed to perform migrations!", err.Error())
			} else if applied == 0 {
				cmd.Println("No migrations could be applied: the provided database is already up to date!")
			} else {
				cmd.Println("Applied", applied, "migrations successfully!")
			}
		},
	}
}

var migrateCmd = MakeMigrateCmd()

func init() {
	RootCmd.AddCommand(migrateCmd)
}
