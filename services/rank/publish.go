package rank

import (
	"encoding/json"

	"github.com/FlorentinDUBOIS/achievements/core/redis"
	"github.com/FlorentinDUBOIS/achievements/models"
)

const (
	channel = "rank"
)

// Publish a rank
func Publish(rank *models.Rank) error {
	client := redis.Connection()
	b, err := json.Marshal(rank)
	if err != nil {
		return err
	}

	return client.
		Publish(channel, string(b)).
		Err()
}
