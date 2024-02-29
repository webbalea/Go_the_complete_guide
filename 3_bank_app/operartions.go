package main

import (
	"bank-app/file_utils"
	"fmt"
)

func withdrawMoney() error {
	var amount int
	fmt.Print("Enter amount to withdraw: ")
	_, err := fmt.Scanf("%d", &amount)
	if err != nil {
		return fmt.Errorf("Invalid input for amount: %w", err)
	}

	if amount <= 0 {
		return fmt.Errorf("Invalid amount. Please enter a positive value.")
	} else if amount > balance {
		return fmt.Errorf("Insufficient funds.")
	}

	newBalance := balance - amount

	err = file_utils.WriteToFile("balance.txt", newBalance)
	if err != nil {
		return err
	}

	fmt.Printf("Withdrew %d. New balance: %d\n", amount, balance)
	return nil
}
