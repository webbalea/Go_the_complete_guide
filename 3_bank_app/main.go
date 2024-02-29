package main

import (
	"bank-app/file_utils"
	"fmt"
)

var balance int = 0

func main() {
	fmt.Println("Hello and welcome to GO Bank!")

	for {
		balance, _ = file_utils.ReadFromFile("balance.txt")

		fmt.Println("Choose an option \n")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		var err error

		fmt.Println("Enter your choice: ")
		_, err = fmt.Scan(&choice) // Read user input

		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			continue // Skip to the next iteration of the loop
		}

		switch choice {
		case 1:
			checkBalance()
		case 2:
			if err := depositMoney(); err != nil {
				fmt.Println(err)
			}
		case 3:
			if err := withdrawMoney(); err != nil {
				fmt.Println(err)
			}
		case 4:
			fmt.Println("Exiting...")
			return // Exit the program directly
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}

func checkBalance() {
	fmt.Printf("Your balance is: %d\n", balance)
}

func depositMoney() error {
	var amount int
	fmt.Print("Enter amount to deposit: ")
	_, err := fmt.Scanf("%d", &amount)
	if err != nil {
		return fmt.Errorf("Invalid input for amount: %w", err)
	}

	if amount <= 0 {
		return fmt.Errorf("Invalid amount. Please enter a positive value.")
	}

	newBalance := balance + amount

	err = file_utils.WriteToFile("balance.txt", newBalance)
	if err != nil {
		return err
	}

	fmt.Printf("Deposited %d. New balance: %d\n", amount, balance)
	return nil
}
