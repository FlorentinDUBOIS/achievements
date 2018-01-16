package cmd

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(databaseUpdateCmd)
}

var databaseUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the database",
	Run: func(cmd *cobra.Command, arguments []string) {
		db := store.Connection()
		defer db.Close()

		user := new(models.User)
		achievement := new(models.Achievement)
		accomplished := new(models.Accomplished)

		tables := []interface{}{user, achievement, accomplished}
		if err := db.AutoMigrate(tables...).Error; err != nil {
			log.Fatal(err)
		}
	},
}
