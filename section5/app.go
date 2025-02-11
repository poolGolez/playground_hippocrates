package main

import (
	"fmt"

	"example.com/bank-account/repository"
)

const BANK_FILE_LOCATION = "./resources/balance.txt"

func main() {

	for {
		choice := promptUserAction()
		switch choice {

		case 1:
			accountBalance := readBalance()
			fmt.Println("=============================================")
			fmt.Printf("Your balance is %.2f\n", accountBalance)
			fmt.Println("=============================================")

		case 2:
			accountBalance := readBalance()
			accountBalance = withdrawAction(accountBalance)
			saveBalance(accountBalance)
			fmt.Println("=============================================")
			fmt.Printf("Your remaining balance is %.2f\n", accountBalance)
			fmt.Println("=============================================")

		case 3:
			accountBalance := readBalance()
			accountBalance = depositAction(accountBalance)
			saveBalance(accountBalance)
			fmt.Println("=============================================")
			fmt.Printf("Your updated balance is %.2f\n", accountBalance)
			fmt.Println("=============================================")

		default:
			fmt.Println("=============================================")
			fmt.Println("Thank you for banking with us!")
			fmt.Println("=============================================")
			return
		}
	}
}

func readBalance() float64 {
	return repository.ReadFromFile(BANK_FILE_LOCATION)
}

func saveBalance(balance float64) {
	repository.WriteToFile(balance, BANK_FILE_LOCATION)
}

func promptUserAction() int {
	var choice int
	welcomeText := `Welcome to the bank! How can I help you?
	1. Check Balance
	2. Withdraw amount
	3. Deposit amount
	4. Exit`
	fmt.Println(welcomeText)

	for choice == 0 || choice > 4 {
		fmt.Print("Choice: ")
		fmt.Scan(&choice)

		if choice > 4 {
			fmt.Println("Invalid choice; please choose again")
		}
	}

	return choice
}

func withdrawAction(accountBalance float64) float64 {
	var amount float64
	fmt.Print("Enter amount to withdraw:")
	fmt.Scan(&amount)

	return accountBalance - amount
}

func depositAction(accountBalance float64) float64 {
	var amount float64
	fmt.Print("Enter amount to deposit:")
	fmt.Scan(&amount)

	return accountBalance + amount
}
