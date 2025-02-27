package loan

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

func FetchAllLoans() []Loan {
	return loans
}
