package main

import (
	"fmt"
)

func main() {
	var revenue, expense, taxRate float64;

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expense: ")
	fmt.Scan(&expense)
	fmt.Print("Tax Rate (%%): ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateFinancials(revenue, expense, taxRate)

	fmt.Printf("Earnings before Tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func calculateFinancials(revenue, expense, taxRate float64) (float64,float64, float64) {
	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	return ebt, profit, ratio
}