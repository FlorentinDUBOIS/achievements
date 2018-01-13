package store

import (
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // driver used by gorm
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	db   *gorm.DB
	once sync.Once
)

func connect() {
	var err error

	host := viper.GetString("database.host")
	port := viper.GetInt("database.port")
	name := viper.GetString("database.name")
	user := viper.GetString("database.user")
	password := viper.GetString("database.password")

	uri := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable", host, port, name, user, password)

	db, err = gorm.Open("postgres", uri)
	if err != nil {
		log.WithFields(log.Fields{
			"error": err,
			"user":  user,
			"host":  host,
			"port":  port,
			"name":  name,
		}).Fatal("Cannot connect to the database")
	}

	db.LogMode(log.GetLevel() == log.DebugLevel)
}

// Connection return the connection to the database
func Connection() *gorm.DB {
	once.Do(connect)

	return db
}
