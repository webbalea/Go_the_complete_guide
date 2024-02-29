package file_utils

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFromFile(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 0, fmt.Errorf("failed to read from file: %w", err)
	}

	balance, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("failed to convert balance to int: %w", err)
	}

	return balance, nil
}

func WriteToFile(filename string, amount int) error {
	// Convert amount to a string
	amountStr := strconv.Itoa(amount)

	// Write the amount string to the specified file
	err := os.WriteFile(filename, []byte(amountStr), 0644)
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}
	return nil
}
