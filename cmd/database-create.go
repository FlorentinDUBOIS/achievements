package cmd

import (
	"github.com/FlorentinDUBOIS/achievements/core/store"
	"github.com/FlorentinDUBOIS/achievements/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(databaseCreateCmd)
}

var databaseCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create the database",
	Run: func(cmd *cobra.Command, arguments []string) {
		db := store.Connection()
		defer db.Close()

		user := new(models.User)
		achievement := new(models.Achievement)

		tables := []interface{}{user, achievement}
		if err := db.CreateTable(tables...).Error; err != nil {
			log.Fatal(err)
		}
	},
}
