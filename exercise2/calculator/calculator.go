package calculator

import "math"

type LoanParameters struct {
	LoanAmount      float64
	InterestPercent float64
	YearsPayable    int64
}

func (params LoanParameters) MonthsPayable() int64 {
	return params.YearsPayable * 12
}

func (params LoanParameters) AnnualInterestRate() float64 {
	return params.InterestPercent / 100
}

func (params LoanParameters) MonthlyInterestRate() float64 {
	return params.AnnualInterestRate() / 12
}

func (params LoanParameters) CalculatePayment() float64 {
	monthlyInterestRate := params.MonthlyInterestRate()
	numerator := monthlyInterestRate * math.Pow(1+monthlyInterestRate, float64(params.MonthsPayable()))
	denominator := math.Pow(1+monthlyInterestRate, float64(params.MonthsPayable())) - 1
	return params.LoanAmount * (numerator / denominator)
}
