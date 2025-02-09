package main

import "fmt"

func main() {
	accountBalance := 1200.0
	
	choice := promptUserAction()
	switch(choice) {
	case 1:
		fmt.Printf("Your balance is %.2f\n", accountBalance)
	case 2:
		accountBalance = withdrawAction(accountBalance)
		fmt.Printf("Your remaining balance is %.2f\n", accountBalance)
	case 3:
		accountBalance = depositAction(accountBalance)
		fmt.Printf("Your updated balance is %.2f\n", accountBalance)
	default:
		fmt.Println("Exiting now.")
	}

	fmt.Println("Thank you for banking with us!")
}

func promptUserAction() int8 {
	var choice int8;
	welcomeText:=`Welcome to the bank! How can I help you?
	1. Check Balance
	2. Withdraw amount
	3. Deposit amount
	4. Exit`
	fmt.Println(welcomeText);
	
	fmt.Print("Choice: ")
	fmt.Scan(&choice)

	return choice
}

func withdrawAction(accountBalance float64) float64 {
	var amount float64
	fmt.Print("Enter amount to withdraw:")
	fmt.Scan(&amount)

	return accountBalance - amount
}

func depositAction(accountBalance float64) float64 {
	var amount float64;
	fmt.Print("Enter amount to deposit:")
	fmt.Scan(&amount)
	
	return accountBalance + amount
}