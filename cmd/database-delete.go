package cmd

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(databaseDeleteCmd)
}

var databaseDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the forge database",
	Run: func(cmd *cobra.Command, arguments []string) {
		db := store.Connection()
		defer db.Close()

		tables := []interface{}{
			new(models.Achievement),
			new(models.User),
		}

		if err := db.DropTable(tables...).Error; err != nil {
			log.Fatal(err)
		}
	},
}
