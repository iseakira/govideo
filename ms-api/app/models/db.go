package models

import (
	"fmt"
	"log"

	"ms-api/config"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var DbConnection *gorm.DB

func connectDb() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		config.Config.DbHost,
		config.Config.DbPort,
		config.Config.DbUser,
			config.Config.DbName,
		config.Config.DbPassword,
		config.Config.DbSslMode,

	)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func init() {
	db,err := connectDb()
	DbConnection = db
	if err != nil {
		log.Fatalln(err)
	}
}