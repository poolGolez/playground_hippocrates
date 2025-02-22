package schedule

import (
	clc "example.com/amortization/calculator"
)

func CalculateSchedule(params clc.LoanParameters) Schedule {
	rows := make([]scheduleRow, params.MonthsPayable())
	payment := params.CalculatePayment()

	for i := 0; i < len(rows); i++ {
		grossBalance := 0.0
		if i == 0 {
			grossBalance = params.LoanAmount
		} else {
			grossBalance = rows[i-1].RemainingBalance
		}
		interest := grossBalance * params.MonthlyInterestRate()
		paidToPrincipal := payment - interest
		rows[i] = scheduleRow{
			PaidToInterest:   interest,
			PaidToPrincipal:  paidToPrincipal,
			RemainingBalance: grossBalance - paidToPrincipal,
		}
	}

	return Schedule{
		Rows: rows,
	}
}

type Schedule struct {
	Rows []scheduleRow
}

type scheduleRow struct {
	PaidToInterest   float64
	PaidToPrincipal  float64
	RemainingBalance float64
}
