package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(databaseCmd)
}

var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Database related commands",
}
