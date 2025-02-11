package utils

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFromFile(location string) float64 {
	data, readError := os.ReadFile(location)
	if readError != nil {
		panic("Unable to retrieve current balance.")
	}

	balance, parseError := strconv.ParseFloat(string(data), 64)
	if parseError != nil {
		panic("Corrupted data detected on file")
	}
	return balance
}

func WriteToFile(balance float64, location string) {
	balanceText := fmt.Sprintf("%.2f", balance)
	error := os.WriteFile(location, []byte(balanceText), 0644)
	if error != nil {
		panic("Error writing to file")
	}

}
