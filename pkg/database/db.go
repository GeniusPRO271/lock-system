package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Start_db() *gorm.DB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=mysecretpassword dbname=postgres port=5430 sslmode=disable",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Drop all tables before migration

	if err := db.AutoMigrate(&User{}, &Device{}, &Whitelist{}, &Space{}, &Log{}, &Role{}); err != nil {
		log.Fatal("Failed migration:", err)
	}

	// Create roles
	// Add table suffix when creating tables
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})

	return db
}
