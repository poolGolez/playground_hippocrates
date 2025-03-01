package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "loans.db")
	if err != nil {
		panic("Unable to connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	sql := `
		CREATE TABLE IF NOT EXISTS loans (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			principal DECIMAL(19, 6) NOT NULL,
			annual_interest_rate DECIMAL(19, 6) NOT NULL,
			years_payable INT NOT NULL
		);
	`
	_, err := DB.Exec(sql)
	if err != nil {
		panic("Could not initialize Loans table.")
	}
}
