package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("env file load failed")
	}

	dsn := os.Getenv("CONN_STR")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	fmt.Println("Database connection was successful")

	loc, _ := time.LoadLocation("America/Buenos_Aires")
	now := time.Now().In(loc)

	fmt.Println(now)
	err = db.QueryRow("SELECT NOW()").Scan(&now)
	if err != nil {
		log.Fatal("failed to execute query", err)
	}

	fmt.Println(now)
}
