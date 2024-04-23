package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start_db() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Drop all tables before migration

	if err := db.AutoMigrate(&User{}, &Device{}, &Space{}, &Whitelist{}, &Log{}, &Role{}); err != nil {
		log.Fatal("Failed migration:", err)
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	return db
}
