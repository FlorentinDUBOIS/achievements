package cmd

import (
	"github.com/FlorentinDUBOIS/achievements/core/pg"
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
		db := pg.Connection()
		defer db.Close()

		user := new(models.User)
		achievement := new(models.Achievement)
		accomplished := new(models.Accomplished)

		tables := []interface{}{user, achievement, accomplished}
		if err := db.CreateTable(tables...).Error; err != nil {
			log.Fatal(err)
		}

		db.
			Model(accomplished).
			AddForeignKey("user_id", "users (id)", "cascade", "cascade")

		if err := db.Error; err != nil {
			log.Fatal(err)
		}

		db.
			Model(accomplished).
			AddForeignKey("achievement_id", "achievements (id)", "cascade", "cascade")

		if err := db.Error; err != nil {
			log.Fatal(err)
		}
	},
}
