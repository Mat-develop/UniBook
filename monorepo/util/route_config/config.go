package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString = ""
	Port             = 0
	SecretKey        []byte
)

func Load() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file, reading config from environment variables")
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	var err error
	Port, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil || Port == 0 {
		Port, err = strconv.Atoi(os.Getenv("API_PORT"))
		if err != nil || Port == 0 {
			Port = 9000
		}
	}

	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, dbName,
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
