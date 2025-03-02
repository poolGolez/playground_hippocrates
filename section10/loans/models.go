package loans

type Loan struct {
	Id                 int64
	Principal          float64 `binding:"required"`
	AnnualInterestRate float64 `binding:"required"`
	YearsPayable       int     `binding:"required"`
}
