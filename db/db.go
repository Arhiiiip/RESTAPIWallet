package db

import (
	model "RESTAPIWallet/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func InitDB() (*gorm.DB, error) {

	godotenv.Load()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASSWORD"), os.Getenv("PSQL_DBNAME"),
		os.Getenv("PSQL_HOST"), os.Getenv("PSQL_PORT"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Wallet{}); err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Transaction{}); err != nil {
		return nil, err
	}

	return db, nil
}
