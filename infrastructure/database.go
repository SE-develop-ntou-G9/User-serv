package infrastructure

import (
	Model "golangAPI/infrastructure/model"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func NewDBInstance() (*gorm.DB, error) {
	godotenv.Load(".env")
	dsn := os.Getenv("DB_URL")

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.AutoMigrate(&Model.UserModel{}, &Model.DriverModel{}, &Model.NotifyModel{})
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	log.Println("Database connected and migrate successfully!")
	return DB, err
}
