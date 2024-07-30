package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBClient struct {
	DB *sql.DB
}

func InitDB() (*DBClient, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	dsn := os.Getenv("CONN_STR")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connection was successful")
	if err := db.Ping(); err != nil {
		return nil, err
	}

	createTaskTable, err := os.ReadFile("data/migrate.sql")
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(string(createTaskTable)); err != nil {
		return nil, err
	}

	return &DBClient{db}, nil
}
