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

func (params LoanParameters) InterestDecimal() float64 {
	return params.InterestPercent / 100
}

func (params LoanParameters) CalculatePayment() float64 {
	monthlyInterestRate := params.InterestDecimal() / 12
	numerator := params.LoanAmount * monthlyInterestRate * math.Pow(1+monthlyInterestRate, float64(params.MonthsPayable()))
	denominator := math.Pow(1+monthlyInterestRate, float64(params.MonthsPayable())) - 1
	return numerator / denominator
}
