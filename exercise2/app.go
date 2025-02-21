package main

import (
	"fmt"

	calc "example.com/amortization/calculator"
	sched "example.com/amortization/schedule"
)

func main() {
	inputs := getInputs()

	schedule := sched.CalculateSchedule(inputs)
	fmt.Println(schedule)

	fmt.Printf("Regular Payment per month: %.2f\n", inputs.CalculatePayment())
}

func getInputs() calc.LoanParameters {
	return calc.LoanParameters{
		LoanAmount:      2500000,
		InterestPercent: 6.75,
		YearsPayable:    10,
	}
}
