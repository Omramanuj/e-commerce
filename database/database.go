package database

import (
	"github.com/Om/models"
	"log"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)

func InitDB() *gorm.DB {
    dbURL := "postgres://om:1234@localhost:5432/postgres"

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(&models.Product{})

    return db
}
