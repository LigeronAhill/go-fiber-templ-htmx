package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func Init() {
	if err := godotenv.Load(); err != nil {
		log.Println("config read warning", "no '.env' file!")
		return
	}
}

type DatabaseConfig struct {
	URL string
}

func NewDatabaseConfig() *DatabaseConfig {
	dv := "postgres://postgres:postgres@localhost:5432/database"
	url := getString("DATABASE_URL", dv)
	return &DatabaseConfig{
		URL: url,
	}
}

type LogConfig struct {
	Level  int
	Format string
}

func NewLogConfig() *LogConfig {
	value := getInt("LOG_LEVEL", 0)
	return &LogConfig{
		Level:  value,
		Format: getString("LOG_FORMAT", "console"),
	}
}

func getString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	} else {
		return value
	}
}

func getInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	n, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return n
}
