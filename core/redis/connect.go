package redis

import (
	"sync"

	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	client *redis.Client
	once   sync.Once
)

func connect() {
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")

	log.WithFields(log.Fields{
		"address": address,
		"db":      db,
	}).Debug("Try to connect to redis")

	client = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.WithFields(log.Fields{
			"address": address,
			"db":      db,
		}).Fatal("Cannot connect to database")
	}
}

// Connection return the redis connection
func Connection() *redis.Client {
	once.Do(connect)

	return client
}
