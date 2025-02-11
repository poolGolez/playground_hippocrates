package repository

import (
	"encoding/csv"
	"os"
	"strconv"

	"example.com/bank-account/domain"
)

const ACCOUNTS_TABLE = "./resources/accounts-table.csv"

func Find(id string) *domain.Account {
	file, err := os.Open(ACCOUNTS_TABLE)
	if err != nil {
		panic("Error opening the file")
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Error reading from file")
	}

	var targetRow []string = nil
	for i, record := range records {
		if record[0] == id {
			targetRow = record
		}
	}

	if targetRow == nil {
		return nil
	}

	balance, _ := strconv.ParseFloat(targetRow[1], 64)

	return &domain.Account{
		Id:      targetRow[0],
		Balance: balance,
	}
}

func Update(account domain.Account) domain.Account {
	return account
}
