package repository

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"example.com/bank-account/domain"
)

const ACCOUNTS_TABLE = "./resources/accounts-table.csv"

func Find(id string) *domain.Account {
	records := getCsvRows()

	var targetRow []string = nil
	for _, record := range records {
		if record[0] == id {
			targetRow = record
		}
	}

	if targetRow == nil {
		return nil
	}

	return toAccount(targetRow)
}

func Update(account domain.Account) domain.Account {
	records := getCsvRows()

	for _, record := range records {
		if record[0] == account.Id {
			record[1] = strconv.FormatFloat(account.Balance, 'f', 2, 64)
		}
	}

	writeCsv(records)

	return account
}

func getCsvRows() [][]string {
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

	return records[1:]
}

func writeCsv(rows [][]string) {
	file, err := os.OpenFile(ACCOUNTS_TABLE, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic("Error opening the file")
	}
	defer file.Close()

	headers := [2]string{"ID", "Balance"}

	writer := csv.NewWriter(file)
	records := [][]string{}
	records = append(records, headers[:])
	records = append(records, rows...)

	fmt.Println(records)
	writer.WriteAll(records)

	if err := writer.Error(); err != nil {
		log.Fatalln("Error saving records in file", err)
		panic("Error saving records in file")
	}
}

func toAccount(row []string) *domain.Account {
	balance, _ := strconv.ParseFloat(row[1], 64)
	return &domain.Account{
		Id:      row[0],
		Balance: balance,
	}
}
