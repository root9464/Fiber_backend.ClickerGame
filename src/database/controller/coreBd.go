package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Data struct {
	*gorm.DB
}

var DB Data

func ConnectToDB() (Data, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки переменных окружения: %v", err)
	}

	dsn := os.Getenv("DB_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных: %v", err)
	}
	DB = Data{db}
	return DB, nil
}
