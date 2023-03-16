package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	// Carga el archivo .env en la aplicación
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error cargando el archivo .env: %v", err)
	}
}

var db *gorm.DB

func InitDB() {
	var err error
	dsn := os.Getenv("DB_DNS")
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func GetDB() *gorm.DB {
	return db
}
