package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5;

func main() {
	var initialInvestment float64;
	var interestRate float64;
	var years int

	fmt.Print("Enter investment value: ")
	fmt.Scan(&initialInvestment)

	fmt.Printf("Enter interest rate (%%): ")
	fmt.Scan(&interestRate)
	
	fmt.Printf("Enter number of years: ")
	fmt.Scan(&years)

	futureValue, inflatedFutureValue := computeFutureValues(initialInvestment,interestRate, years)

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Adjusted with %.2f%% inflation: %.2f\n", inflationRate, inflatedFutureValue)
}

func computeFutureValues(initialInvestment float64, interestRate float64, years int) (float64, float64) {
	futureValue := initialInvestment * math.Pow(1 + interestRate/100, float64(years))
	inflatedFutureValue := futureValue / math.Pow(1 + inflationRate/100, float64(years))

	return futureValue, inflatedFutureValue
}