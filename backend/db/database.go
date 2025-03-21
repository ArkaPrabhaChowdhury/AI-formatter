package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDatabase() error {
	_ = godotenv.Load()

	fmt.Println("Connecting to database...")

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = sql.Open("pgx", dsn)
	if err != nil {
		fmt.Println("Failed to connect to database:", err)
		return err
	}

	fmt.Println("Pinging database...")
	return DB.Ping()
}
