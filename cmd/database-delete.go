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

		user := new(models.User)
		achievement := new(models.Achievement)
		accomplished := new(models.Accomplished)

		db.
			Model(accomplished).
			RemoveIndex("accomplisheds_user_id_users_id_foreign")

		if err := db.Error; err != nil {
			log.Fatal(err)
		}

		db.
			Model(accomplished).
			RemoveIndex("accomplisheds_achievement_id_achievements_id_foreign")

		if err := db.Error; err != nil {
			log.Fatal(err)
		}

		tables := []interface{}{user, achievement, accomplished}
		if err := db.DropTable(tables...).Error; err != nil {
			log.Fatal(err)
		}
	},
}
