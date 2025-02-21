package main

import (
	"fmt"
	"math"
)

func main() {
	/* INPUTS
	* Loan amount
	* Interest Rate
	* Number of Years
	* Monthly Payment
	 */
	inputs := getInputs()

	interestDecimal := (inputs.InterestPercent / 100) / 12
	common := (1 + interestDecimal)
	months := inputs.YearsPayable * 12
	numerator := inputs.LoanAmount * interestDecimal * math.Pow(common, float64(months))
	denominator := math.Pow(common, float64(months)) - 1

	payment := numerator / denominator
	// fmt.Printf("Common: %.6f\n", common)
	// fmt.Printf("Interest decimal: %.6f\n", interestDecimal)
	// fmt.Printf("Months: %.2d\n", months)

	// fmt.Printf("Numerator: %.2f\n", numerator)
	// fmt.Printf("Denominator: %.2f\n", denominator)
	fmt.Printf("Regular Payment per month: %.2f\n", payment)

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

		interest := remainingBalance * interestDecimal
		paidToPrincipal := payment - interest
		schedule[month][0] = interest
		schedule[month][1] = paidToPrincipal
		schedule[month][2] = remainingBalance - paidToPrincipal
	}

	for index, record := range schedule {
		fmt.Printf("Month %d: %.2f\n", index+1, record)
	}
}

func getInputs() LoanParameters {
	return LoanParameters{
		LoanAmount:      2500000,
		InterestPercent: 6.8,
		YearsPayable:    5,
	}
}

type LoanParameters struct {
	LoanAmount      float64
	InterestPercent float64
	YearsPayable    int64
}
