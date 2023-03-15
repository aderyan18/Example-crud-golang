package config

import (
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("")
	database := os.Getenv("DB_NAME")

	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"

	config := gorm.Config{}

	db, err := gorm.Open(mysql.Open(dsn), &config)
	if err != nil {
		panic("Failed to connect to database!" + os.Getenv("DB_USER"))
	}

	DB = db
}

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from database")
	}
	dbSQL.Close()
}
