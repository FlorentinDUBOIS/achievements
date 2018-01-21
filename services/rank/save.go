package rank

import (
	"github.com/FlorentinDUBOIS/achievements/core/redis"
	"github.com/FlorentinDUBOIS/achievements/models"
	"github.com/FlorentinDUBOIS/achievements/services/utils"
)

const (
	accomplisheds = "accomplisheds"
)

// Save a rank
func Save(rank *models.Rank) error {
	client := redis.Connection()
	if err := client.Set(rank.UserID, rank.Count, 0).Err(); err != nil {
		return err
	}

	uids, err := client.LRange(accomplisheds, 0, -1).Result()
	if err != nil {
		return err
	}

	if !utils.Contains(uids, rank.UserID) {
		if err := client.RPush(accomplisheds, rank.UserID).Err(); err != nil {
			return err
		}
	}

	return nil
}
