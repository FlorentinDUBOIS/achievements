package cmd

import (
	"time"

	"github.com/FlorentinDUBOIS/achievements/core/pg"
	"github.com/FlorentinDUBOIS/achievements/models"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	databaseCmd.AddCommand(databaseInsertCmd)
}

var achievements = []models.Achievement{
	{
		ID:          "f4ecf79c-b218-400b-8f90-c239f2c118cd",
		Name:        "Le début de la gloire !",
		Description: "Se connecter à l'application",
	},
	{
		ID:          "555210a2-f862-4ee2-a2da-0fb4ad2e167a",
		Name:        "Se faire un ami",
		Description: "Scanner le badge d'une personne",
	},
	{
		ID:          "06b6164d-e702-462c-b374-c1a66e13462b",
		Name:        "Se faire 5 ami(e)s",
		Description: "Scanner le badge de 5 personnes",
	},
	{
		ID:          "4b22a05f-6af4-4ec7-aa7a-a4aab223e3e9",
		Name:        "Se faire 10 ami(e)s",
		Description: "Scanner le badge de 10 personnes",
	},
	{
		ID:          "c24e40d1-4c9b-44c4-843e-5f6ff6e8a535",
		Name:        "Se faire 15 ami(e)s",
		Description: "Scanner le badge de 15 personnes",
	},
	{
		ID:          "aa7afd00-5845-46c0-84c5-91e0387744eb",
		Name:        "Se faire 20 ami(e)s",
		Description: "Scanner le badge de 20 personnes",
	},
	{
		ID:          "47a87d28-dcd4-45d4-bc88-621419350cda",
		Name:        "Se faire 25 ami(e)s",
		Description: "Scanner le badge de 25 personnes",
	},
	{
		ID:          "a70b502b-7e4d-4d87-9d40-55d3fabe10a2",
		Name:        "Se faire 30 ami(e)s",
		Description: "Scanner le badge de 30 personnes",
	},
	{
		ID:          "c0b22e94-9a77-47d6-9697-8eb8efb78a03",
		Name:        "Se faire 35 ami(e)s",
		Description: "Scanner le badge de 35 personnes",
	},
	{
		ID:          "2115c95d-dc6c-4288-8813-3bb6a1dbc5ef",
		Name:        "Se faire 40 ami(e)s",
		Description: "Scanner le badge de 40 personnes",
	},
	{
		ID:          "0aabf13d-e61c-4add-99ed-2d9299410b91",
		Name:        "Se faire 45 ami(e)s",
		Description: "Scanner le badge de 45 personnes",
	},
	{
		ID:          "44d0888c-769f-42bc-bad8-0574c615119d",
		Name:        "Se faire 50 ami(e)s",
		Description: "Scanner le badge de 50 personnes",
	},
	{
		ID:          "5d7af3bd-067c-4186-b508-7d6387af47fa",
		Name:        "Se faire 55 ami(e)s",
		Description: "Scanner le badge de 55 personnes",
	},
	{
		ID:          "01180159-58bf-4888-94f7-7b97d81a3390",
		Name:        "Se faire 60 ami(e)s",
		Description: "Scanner le badge de 60 personnes",
	},
	{
		ID:          "bdb278dd-ac3d-4d7c-ac11-652900d612af",
		Name:        "Se faire 65 ami(e)s",
		Description: "Scanner le badge de 65 personnes",
	},
	{
		ID:          "7e2a4465-21ba-4530-a0fc-725cd986d4a9",
		Name:        "Se faire 70 ami(e)s",
		Description: "Scanner le badge de 70 personnes",
	},
	{
		ID:          "e699fab6-3215-4df9-9d9a-1287b877d171",
		Name:        "Se faire 75 ami(e)s",
		Description: "Scanner le badge de 75 personnes",
	},
	{
		ID:          "3e71477a-87a3-4524-9ce7-edaabfbf0ff9",
		Name:        "Se faire 80 ami(e)s",
		Description: "Scanner le badge de 80 personnes",
	},
	{
		ID:          "c68be29f-8282-448d-b1b9-99840eaaffd5",
		Name:        "Vous êtes populaire",
		Description: "S'être fait 80 ami(e)s",
	},
	{
		ID:          "8e2edd9f-4555-4ce8-a825-c0787776c6ad",
		Name:        "Bananonina",
		Description: "Faire un calin à un minion",
	},
	{
		ID:          "d2c9ed30-e029-4932-94f3-2f34052e838c",
		Name:        "Sae",
		Description: "Faire un calin à 2 minions",
	},
	{
		ID:          "bc0b32a7-ba9c-4971-8797-a89957392ac7",
		Name:        "Me want banana",
		Description: "Faire un calin à 3 minions",
	},
	{
		ID:          "226ad8ac-ed94-491f-b604-29dafaab2eb6",
		Name:        "Dul",
		Description: "Faire un calin à 4 minions",
	},
	{
		ID:          "5ab821e2-6d81-4810-8de7-303b8b578baa",
		Name:        "Buttom",
		Description: "Faire un calin à 5 minions",
	},
	{
		ID:          "6de7744d-349f-4412-8427-48efcca5815a",
		Name:        "Sa la ka!",
		Description: "Faire un calin à 6 minions",
	},
	{
		ID:          "b7b4777d-990f-4534-a58d-0773b01ba56b",
		Name:        "Kampai",
		Description: "Faire un calin à 7 minions",
	},
	{
		ID:          "64b94bb5-7faf-4f32-9228-73bb06a7c8bd",
		Name:        "Muak Muak Muak",
		Description: "Faire un calin à 8 minions",
	},
	{
		ID:          "f8eba67f-dfdc-4ff7-9039-c89320e456b1",
		Name:        "Po ka",
		Description: "Faire un calin à 9 minions",
	},
	{
		ID:          "c6bf228e-82b9-478a-952b-56556b712cde",
		Name:        "Underwear...",
		Description: "Faire un calin à 10 minions",
	},
	{
		ID:          "7e2abe37-83bc-4a35-bf0e-e9b76c6c0be5",
		Name:        "F Society",
		Description: "Débloquer tous les succès en moins de 10 secondes",
	},
	{
		ID:          "3c85cd6c-2202-4a05-b540-05e49f3af988",
		Name:        "Je te suis",
		Description: "Suivre le compte twitter du Startup Weekend, à vous de le trouver !",
	},
}

var databaseInsertCmd = &cobra.Command{
	Use:   "insert",
	Short: "Insert data in database",
	Run: func(cmd *cobra.Command, args []string) {
		db := pg.Connection()
		defer db.Close()

		now := time.Now()
		for _, achievement := range achievements {
			achievement.CreatedAt = now
			achievement.UpdatedAt = now

			if err := db.Create(&achievement).Error; err != nil {
				log.WithFields(log.Fields{
					"error":       err,
					"achievement": achievement,
				}).Error("Cannot insert achievement")
			}
		}

	},
}
