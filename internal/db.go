package internal

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"github.com/joho/godotenv"
)

func InitDB() *sql.DB {
	_ = godotenv.Load()
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}
	return db
}
