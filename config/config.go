package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

var once sync.Once
var instance *Config

func getConfig() *Config {
	once.Do(func() {
		username := GetEnv("DB_USERNAME", "postgres")
		password := GetEnv("DB_PASSWORD", "postgres")
		host := GetEnv("DB_HOST", "postgres")
		port := GetEnv("DB_PORT", "5432")
		db_name := GetEnv("DB_NAME", "postgres")
		server_port := GetEnv("SERVER_PORT", "8080")

		dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, db_name)
		instance = &Config{
			DatabaseURL: dsn,
			ServerPort:  server_port,
		}
	})

	return instance
}

func GetEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		Logger.Info("No .env file was found")
	}
	Logger.Info("Initialized Enviroment")
}
