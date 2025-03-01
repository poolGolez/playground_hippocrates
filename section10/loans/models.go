package loans

type Loan struct {
	Id                 int64
	Principal          float64 `json:"principal" binding:"required"`
	AnnualInterestRate float64 `json:"annualInterestRate" binding:"required"`
	YearsPayable       int     `json:"yearsPayable" binding:"required"`
}
