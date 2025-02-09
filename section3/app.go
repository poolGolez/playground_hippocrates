package main

import "fmt"

func main() {
	accountBalance := 1200.0
	var choice int8;

	welcomeText:=`Welcome to the bank! How can I help you?
	1. Check Balance
	2. Withdraw amount
	3. Deposit amount
	4. Exit`
	fmt.Println(welcomeText);
	
	fmt.Print("Choice: ")
	fmt.Scan(&choice)

	switch(choice) {
	case 1:
		fmt.Printf("Your balance is %.2f\n", accountBalance)
	case 2:
		var amount float64
		fmt.Print("Enter amount to withdraw:")
		fmt.Scan(&amount)
		accountBalance -= amount
		fmt.Printf("Your remaining balance is %.2f\n", accountBalance)
	case 3:
		var amount float64;
		fmt.Print("Enter amount to deposit:")
		fmt.Scan(&amount)
		accountBalance += amount
		fmt.Printf("Your updated balance is %.2f\n", accountBalance)
	default:
		fmt.Println("Exiting now.")
	}

}