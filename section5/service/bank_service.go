package service

import (
	"example.com/bank-account/domain"
	"example.com/bank-account/repository"
)

func FetchAccount(id string) domain.Account {
	return repository.Find(id)
}

func Withdraw(account domain.Account, amount float64) domain.Account {
	account.Balance -= amount

	return repository.Update(account)
}

func Deposit(account domain.Account, amount float64) domain.Account {
	account.Balance += amount

	return repository.Update(account)
}
