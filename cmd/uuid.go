package cmd

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(uuidCmd)
}

var uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "generate an uuid",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("uuid: %s\n", uuid.NewV4().String())
	},
}
