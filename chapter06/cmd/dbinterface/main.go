package main

import (
	//_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/internal/database"
	"github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/chapter06/internal/dbinterface"
)

func main() {
	db, err := database.Setup()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	// this wont do anything if commit is successful
	defer tx.Rollback()

	if err := dbinterface.Exec(tx); err != nil {
		panic(err)
	}
	if err := tx.Commit(); err != nil {
		panic(err)
	}
}
