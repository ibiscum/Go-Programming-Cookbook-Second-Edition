package main

import (
	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/database"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	if err := database.Exec(db); err != nil {
		panic(err)
	}
}
