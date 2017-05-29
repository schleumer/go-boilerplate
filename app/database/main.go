package database

import (
	"github.com/jinzhu/gorm"
	"log"
	"../models"
)

var LocalDb *gorm.DB

func BootDb() {
	localDb, err := gorm.Open("sqlite3", "./local.db")

	localDb.AutoMigrate(&models.User{})

	firstUserExists := 0

	localDb.Table("users").Where("id = 1").Count(&firstUserExists)

	if firstUserExists < 1 {
		// senha@123
		localDb.Exec("INSERT INTO users (id, name, username, password, is_system) VALUES (1, 'Admin', 'admin', '$2a$12$IAA9MkM2z.HmneercK.UA.Kbv1SF1xfnmu1Ho7px.0Tdh6UBGqLFS', 1)")
	}

	if err != nil {
		log.Fatal(err)
	}

	LocalDb = localDb
}
