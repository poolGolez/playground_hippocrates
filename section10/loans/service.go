package loans

var loans []Loan = []Loan{
	{
		Principal:          2500000.0,
		AnnualInterestRate: 0.0675,
		YearsPayable:       10,
	},
	{
		Principal:          2500000.0,
		AnnualInterestRate: 0.0675,
		YearsPayable:       15,
	},
}

func FetchAll() []Loan {
	return findAll()
}

func Save(loan *Loan) *Loan {
	loans = append(loans, *loan)
	save(loan)
	return loan
}
