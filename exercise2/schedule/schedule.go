package schedule

import (
	clc "example.com/amortization/calculator"
)

func CalculateSchedule(params clc.LoanParameters) Schedule {
	rows := make([]ScheduleRow, params.MonthsPayable())
	payment := params.CalculatePayment()

	for i := 0; i < len(rows); i++ {
		grossBalance := 0.0
		if i == 0 {
			grossBalance = params.LoanAmount
		} else {
			grossBalance = rows[i-1].RemainingBalance
		}
		interest := grossBalance * params.MonthlyInterestRate()

		rows[i] = ScheduleRow{
			PaidToInterest:   interest,
			PaidToPrincipal:  payment - interest,
			RemainingBalance: grossBalance - payment,
		}
	}

	return Schedule{
		rows: rows,
	}
}

type Schedule struct {
	rows []ScheduleRow
}

type ScheduleRow struct {
	PaidToInterest   float64
	PaidToPrincipal  float64
	RemainingBalance float64
}
