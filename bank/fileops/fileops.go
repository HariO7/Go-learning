package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteBalance(fileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func ReadBalance(fileName string) (float64, error) {
	balance, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, err
	}
	balanceValue, err := strconv.ParseFloat(string(balance), 64)

	if err != nil {
		return 1000, err
	}

	return balanceValue, nil
}
