package database

import (
	"fmt"
	"irctc/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := os.Getenv("dsn")
	if dsn == "" {
		log.Panic("No connection string found")
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("Failed to connect to the database!")
	}

	err = db.AutoMigrate(&models.User{}, &models.Train{}, &models.Booking{})
	if err != nil {
		panic("Failed to run migrations!")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		panic("Failed to ping the database!")
	}

	DB = db
	fmt.Println("Connected to the database successfully!")
}

func GetDB() *gorm.DB {
	return DB
}
