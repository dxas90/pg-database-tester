package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
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

func isMongoURI(uri string) bool {
	return strings.HasPrefix(uri, "mongodb://") || strings.HasPrefix(uri, "mongodb+srv://")
}

func connectMongo(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(ctx)

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal("failed to ping MongoDB:", err)
	}

	fmt.Println("Successfully connected to MongoDB")
}

func connectPostgres(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}

	var now time.Time
	db.Raw("SELECT NOW()").Scan(&now)

	fmt.Println("Current time from DB:", now)
}

func main() {
	dsn := os.Getenv("DATABASE_URL")

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

	if isMongoURI(dsn) {
		connectMongo(dsn)
	} else {
		connectPostgres(dsn)
	}
}
