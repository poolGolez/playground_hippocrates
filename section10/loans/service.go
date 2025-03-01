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
