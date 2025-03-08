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
	if err := createUsersTable(); err != nil {
		panic("Could not initialize Users table.")
	}

	if err := createLoansTable(); err != nil {
		panic("Could not initialize Loans table.")
	}
}

func createLoansTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS loans (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			principal DECIMAL(19, 6) NOT NULL,
			annual_interest_rate DECIMAL(19, 6) NOT NULL,
			years_payable INT NOT NULL,
			user_id INTEGER NOT NULL REFERENCES users(id)
		);
	`
	_, err := DB.Exec(sql)
	return err
}

func createUsersTable() error {
	sql := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			date_registered TIMESTAMPTZ NOT NULL
		);
	`
	_, err := DB.Exec(sql)
	return err
}
