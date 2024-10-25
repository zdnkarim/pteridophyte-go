package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	var err error

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if err != nil {
		log.Println(err.Error())
	}

	// Generate Connection String
	// conTemplate := "host=%s port=%d dbname=%s user=%s password=%s sslmode=disable"

	// Another template
	// conTemplate := "postgresql://%s:%s@%s:%d/%s?sslmode=disable"
	conTemplate := "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable"
	//user:password@host:port/dbname

	connStr := fmt.Sprintf(
		conTemplate,
		host,
		port,
		user,
		password,
		dbname,
	)

	// Connect to database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Println("Cannot connect to database")
	}

	return db
}
