package db

import (
	"log"
	"smc-wallet-be/internal/module/user/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	dbURL := "postgres://root:secret@localhost:5432/golang"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&entities.User{})

	return db
}
