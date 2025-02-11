package repository

import "example.com/bank-account/domain"

const BANK_FILE_LOCATION = "./resources/balance.txt"

func Find(id string) domain.Account {

	return domain.Account{
		Id:      "12345",
		Balance: 1200.00,
	}
}
