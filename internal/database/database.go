package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	DBname   string
	Username string
	Port     string
	Password string
	Sslmode  string
	TimeZone string
}

func ConnectToDB(cfg Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", cfg.Host, cfg.Username, cfg.Password, cfg.DBname, cfg.Port, cfg.Sslmode, cfg.TimeZone)

	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func CloseDB(db *gorm.DB) {
	if sqlDBobj, err := db.DB(); err != nil {
		log.Fatalf("error in getting sqlDB object: %s", err.Error())
		return
	} else {
		if err = sqlDBobj.Close(); err != nil {
			log.Fatalf("error in closing DB: %s", err.Error())
		}
	}
}
