package db

import (
	"blitztracker_api/src/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	// load config
	env, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	dbURL := env.POSTGRES_CONN

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	//migrations goes here

	return db
}
