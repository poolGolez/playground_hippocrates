package main

import (
	"fmt"

	calc "example.com/amortization/calculator"
	renderer "example.com/amortization/renderer"
	sched "example.com/amortization/schedule"
)

func main() {
	inputs := getInputs()

	schedule := sched.CalculateSchedule(inputs)
	fmt.Println(renderer.Format(schedule))

	fmt.Printf("Regular Payment per month: %s\n", renderer.FormatMoney(inputs.CalculatePayment()))
	fmt.Printf("Total Loan Interest: %s\n", renderer.FormatMoney(inputs.CalculateTotalInterest()))
	fmt.Printf("Interest to Principal Ratio: %s\n", renderer.FormatMoney(inputs.CalculateInterestToPrincipalRatio()))
}

func getInputs() calc.LoanParameters {
	return calc.LoanParameters{
		LoanAmount:      2500000,
		InterestPercent: 6.75,
		YearsPayable:    10,
	}
}
