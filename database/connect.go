package database

import (
	"fmt"
	"strconv"

	"fiber.osaxon/config"
	"fiber.osaxon/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() {
	var err error

	p := config.Config("DB_PORT")
	port, err := strconv.Atoi(p)

	if err != nil {
		panic("Failed to connect database")
	}

	dsn := fmt.Sprintf(
		"host=localhost port=%d user=%s password=%s dbname=%s sslmode=disable",
		port,
		config.Config("DB_USER"),
		config.Config("DB_PASSWORD"),
		config.Config("DB_NAME"),
	)

	println(dsn)

	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.Product{})
	fmt.Println("Database Migrated")
}
