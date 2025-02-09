package main

import "fmt"

func main() {
	var revenue, expense, taxRate float64;

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Expense: ")
	fmt.Scan(&expense)
	fmt.Print("Tax Rate (%%): ")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	profit := ebt * (1 - taxRate/100)
	ratio := ebt/profit

	fmt.Printf("Earnings before Tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}