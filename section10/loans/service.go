package loans

func FetchAll() []Loan {
	return findAll()
}

func FetchById(id int64) *Loan {
	return find(id)
}

func Save(loan *Loan) *Loan {
	save(loan)

	return loan
}

func Update(loan *Loan, params *UpdateLoanParams) *Loan {
	loan.Principal = params.Principal
	loan.AnnualInterestRate = params.AnnualInterestRate
	loan.YearsPayable = params.YearsPayable

	update(loan)
	return loan
}

func Delete(loan *Loan) {
	delete(loan)
}

type UpdateLoanParams struct {
	Principal          float64 `binding:"required"`
	AnnualInterestRate float64 `binding:"required"`
	YearsPayable       int     `binding:"required"`
}
