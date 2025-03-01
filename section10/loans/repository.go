package loans

import (
	"example.com/gin/loaney/db"
)

func save(loan *Loan) error {
	sql := `
	INSERT INTO loans (principal, annual_interest_rate, years_payable)
	VALUES(?, ?, ?);
	`

	result, err := db.DB.Exec(sql, loan.Principal, loan.AnnualInterestRate, loan.YearsPayable)
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

	rows, err := db.DB.Query(sql)
	if err != nil {
		panic("Failed querying loans.")
	}
	defer rows.Close()

	var results = make([]Loan, 0)
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
	result := db.DB.QueryRow(sql, id)

	var loan Loan
	err := result.Scan(&loan.Id, &loan.Principal, &loan.AnnualInterestRate, &loan.YearsPayable)
	if err != nil {
		return nil
	}
	return &loan
}
