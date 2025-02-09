package main

import "fmt"

func main() {
	accountBalance := 1200.0

	for {
		choice := promptUserAction()
		switch choice {
		case 1:
			fmt.Println("=============================================")
			fmt.Printf("Your balance is %.2f\n", accountBalance)
			fmt.Println("=============================================")
		case 2:
			accountBalance = withdrawAction(accountBalance)
			fmt.Println("=============================================")

			fmt.Printf("Your remaining balance is %.2f\n", accountBalance)
			fmt.Println("=============================================")
		case 3:
			accountBalance = depositAction(accountBalance)
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
