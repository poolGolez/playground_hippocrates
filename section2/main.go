package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5;
	var initialInvestment float64;
	var interestRate float64;
	var years int16

	fmt.Print("Enter investment value: ")
	fmt.Scan(&initialInvestment)

	fmt.Printf("Enter interest rate (%%): ")
	fmt.Scan(&interestRate)
	
	fmt.Printf("Enter number of years: ")
	fmt.Scan(&years)

	var futureValue = initialInvestment * math.Pow(1 + interestRate/100, float64(years))
	var inflatedFutureValue = futureValue / math.Pow(1+inflationRate/100,float64(years))

	fmt.Printf("Future value: %.2f\n", futureValue)
	fmt.Printf("Adjusted with %.2f%% inflation: %.2f\n", inflationRate, inflatedFutureValue)
}