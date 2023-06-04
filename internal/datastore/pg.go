package datastore

import (
	"fmt"
	"github.com/cubbit/cbs3/internal/domain"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQL() *gorm.DB {

	connectString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s %s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
		"sslmode=disable",
	)
	db, err := gorm.Open(postgres.Open(connectString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(&domain.Bucket{},&domain.Object{})
	db.AutoMigrate(&domain.Object{})

	return db
}
