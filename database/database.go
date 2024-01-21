package database

import (
	"fmt"
	"log"
	"os"

	"github.com/AnthonyBrancato/golang-restapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBInstance struct {
	DB *gorm.DB
}

var DB DBInstance

func ConnectDb() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=false TimeZone=Europe/Paris", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DBNAME"))
	db, error := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if error != nil {
		log.Fatal("Failed to connnect to Postgres instance. \n", error)
		os.Exit(2)
	}

	log.Println("Connected.")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("Running migrations")
	db.AutoMigrate(&models.Tickets{})

	DB = DBInstance{
		DB: db,
	}
}
