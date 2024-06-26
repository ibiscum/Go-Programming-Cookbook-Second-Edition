package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Exec takes a new connection
// creates tables, and later drops them
// and issues some queries
func Exec(db *sql.DB) error {
	fmt.Println("Exec()")

	// uncaught error on cleanup, but we always want to cleanup
	defer db.Exec("DROP TABLE example") //nolint

	if err := Create(db); err != nil {
		return err
	}

	if err := Query(db, "Aaron"); err != nil {
		return err
	}
	return nil
}
