package loans

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

func save(loan *Loan) error {
	sql := `
	INSERT INTO loans (principal, annual_interest_rate, years_payable)
	VALUES(?, ?, ?);
	`

	result, err := DB.Exec(sql, loan.Principal, loan.AnnualInterestRate, loan.YearsPayable)
	if err != nil {
		return err
	}

	loan.Id, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func findAll() []Loan {
	sql := `
	SELECT id, principal, annual_interest_rate, years_payable
	FROM loans; 
	`

	rows, err := DB.Query(sql)
	if err != nil {
		panic("Failed querying loans.")
	}
	defer rows.Close()

	var results []Loan
	for rows.Next() {
		var loan Loan
		err = rows.Scan(&loan.Id, &loan.Principal, &loan.AnnualInterestRate, &loan.YearsPayable)
		if err != nil {
			panic("Failed querying loans.")
		}

		results = append(results, loan)
	}

	return results
}

func find(id int64) *Loan {
	sql := `
	SELECT id, principal, annual_interest_rate, years_payable
	FROM loans
	WHERE id = ?;
	`
	result := DB.QueryRow(sql, id)

	var loan Loan
	result.Scan(&loan.Id, &loan.Principal, &loan.AnnualInterestRate, &loan.YearsPayable)
	return &loan
}
