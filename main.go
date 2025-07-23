package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// getEnv returns the value of an environment variable or a default if it's not set
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func main() {
	// Try to get the full DSN first
	dsn := os.Getenv("DATABASE_URL")

	// If PG_DSN is not set, construct it from individual parts
	if dsn == "" {
		user := getEnv("DB_USER", "postgres")
		password := getEnv("DB_PASSWORD", "password")
		host := getEnv("DB_HOST", "localhost")
		port := getEnv("DB_PORT", "5432")
		database := getEnv("DB_NAME", "postgres")
		sslmode := getEnv("DB_SSLMODE", "disable")

		dsn = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
											user, password, host, port, database, sslmode)
	}

	// Connect to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	// Query current time
	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	fmt.Println("Current time from DB:", now)
}
