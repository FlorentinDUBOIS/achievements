package cmd

import (
	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(databaseDeleteCmd)
}

var databaseDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete data in database",
	Run: func(cmd *cobra.Command, args []string) {
		db := pg.Connection()
		defer db.Close()

		if db.Unscoped().Delete(new(models.Achievement)).Error != nil {
			log.WithFields(log.Fields{
				"error": db.Error,
			}).Error("Cannot delete all achievements")
		}
	},
}
