package main

import (
	"fmt"

	"example.com/bank-account/service"
)

const BANK_FILE_LOCATION = "./resources/balance.txt"

func main() {
	for {
		choice := promptUserAction()

		// TODO: Check if not found
		account := service.FetchAccount("12345")

		switch choice {

		case 1:
			accountBalance := account.Balance
			fmt.Println("=============================================")
			fmt.Printf("Your balance is %.2f\n", accountBalance)
			fmt.Println("=============================================")

		case 2:
			var amount float64
			fmt.Print("Enter amount to withdraw:")
			fmt.Scan(&amount)

			updatedAccount := service.Withdraw(*account, amount)
			fmt.Println("=============================================")
			fmt.Printf("Your remaining balance is %.2f\n", updatedAccount.Balance)
			fmt.Println("=============================================")

		case 3:
			var amount float64
			fmt.Print("Enter amount to deposit:")
			fmt.Scan(&amount)

			updatedAccount := service.Deposit(*account, amount)
			fmt.Println("=============================================")
			fmt.Printf("Your updated balance is %.2f\n", updatedAccount.Balance)
			fmt.Println("=============================================")

		default:
			fmt.Println("=============================================")
			fmt.Println("Thank you for banking with us!")
			fmt.Println("=============================================")
			return
		}
	}
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
