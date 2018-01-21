package rank

import (
	"strconv"

	"github.com/FlorentinDUBOIS/achievements/core/redis"
	"github.com/FlorentinDUBOIS/achievements/models"
	log "github.com/sirupsen/logrus"
)

// Get all rank
func Get() ([]models.Rank, error) {
	ranks := make([]models.Rank, 0)
	client := redis.Connection()

	uids, err := client.LRange(accomplisheds, 0, -1).Result()
	if err != nil {
		return nil, err
	}

	for _, uid := range uids {
		item, err := client.Get(uid).Result()
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"uid":   uid,
			}).Error("Cannot retrieve in redis the rank")
			continue
		}

		count, err := strconv.Atoi(item)
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
				"uid":   uid,
				"item":  item,
			}).Error("Cannot convert item to a count")
			continue
		}

		ranks = append(ranks, models.Rank{
			UserID: uid,
			Count:  count,
		})
	}

	return ranks, nil
}
