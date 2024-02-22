package main

import (
	_ "github.com/go-sql-driver/mysql" // we import supported libraries for database/sql
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	// Find .env file
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
