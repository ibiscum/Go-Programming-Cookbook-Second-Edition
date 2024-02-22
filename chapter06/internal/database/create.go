package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //we import supported libraries for database/sql
)

// Create makes a table called example
// and populates it
func Create(db *sql.DB) error {
	fmt.Println("Create()")
	// initialize database
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS example (name VARCHAR(20), created DATETIME)"); err != nil {
		return err
	}

	if _, err := db.Exec(`INSERT INTO example (name, created) values ("Aaron", NOW())`); err != nil {
		return err
	}

	return nil
}
