package database

import (
    "fmt"
	"github.com/Om/models"
	"log"
	"gorm.io/driver/postgres"
    "gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "om"
	password = "1234"
	dbName   = "om"
)

func InitDB() *gorm.DB {
    dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

    db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

    if err != nil {
        log.Fatalln(err)
    }

    db.AutoMigrate(
        &models.User{},
        &models.Order{},
        &models.OrderItem{},
        &models.Category{},
        &models.Product{},
    )

    return db
}
