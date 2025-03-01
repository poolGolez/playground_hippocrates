package loans

func FetchAll() []Loan {
	return findAll()
}

func Save(loan *Loan) *Loan {
	save(loan)
	return loan
}
