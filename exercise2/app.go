package main

import (
	"fmt"

	calc "example.com/amortization/calculator"
)

func main() {
	inputs := getInputs()
	months := inputs.MonthsPayable()
	fmt.Printf("Regular Payment per month: %.2f\n", inputs.CalculatePayment())

	schedule := make([][3]float64, months)
	/**
	* Column 1: Interest
	* Column 2: Paid to principal
	* Column 3: Remaining balance
	 */
	for month := range months {
		var remainingBalance float64
		if month == 0 {
			remainingBalance = inputs.LoanAmount
		} else {
			remainingBalance = schedule[month-1][2]
		}

		interest := remainingBalance * inputs.InterestDecimal()
		paidToPrincipal := inputs.CalculatePayment() - interest
		schedule[month][0] = interest
		schedule[month][1] = paidToPrincipal
		schedule[month][2] = remainingBalance - paidToPrincipal
	}

	for index, record := range schedule {
		fmt.Printf("Month %d: %.2f\n", index+1, record)
	}
}

func getInputs() calc.LoanParameters {
	return calc.LoanParameters{
		LoanAmount:      2500000,
		InterestPercent: 6.8,
		YearsPayable:    1,
	}
}
