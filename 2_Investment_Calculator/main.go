package main

import (
	"fmt"
	"github.com/fatih/color"
	"math"
	"strconv"
)

// GetFloatInput prompts the user for a float value and returns it along with any error.
func GetFloatInput(prompt string) (float64, error) {
	yellow := color.New(color.FgHiYellow)     // Use just yellow color
	fmt.Println(yellow.Sprint(prompt) + ": ") // Add a colon after the prompt
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		return 0, fmt.Errorf("invalid input: %w", err)
	}
	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %w", err)
	}
	return f, nil
}

// Investment represents an investment with its properties.
type Investment struct {
	Amount        float64
	ExpectedRate  float64
	Years         float64
	InflationRate float64
	TaxRate       float64
}

// NewInvestment creates a new Investment object with default values for inflation and tax rates.
func NewInvestment() *Investment {
	return &Investment{
		InflationRate: 2.5, // Default inflation rate
		TaxRate:       20,  // Default tax rate
	}
}

// CalculateFutureValue calculates the future value of the investment.
func (investment *Investment) CalculateFutureValue() float64 {
	return investment.Amount * math.Pow(1+investment.ExpectedRate/100, investment.Years)
}

// CalculateFutureRealValue calculates the future real value considering inflation.
func (investment *Investment) CalculateFutureRealValue() float64 {
	futureValue := investment.CalculateFutureValue()
	return futureValue / math.Pow(1+investment.InflationRate/100, investment.Years)
}

// CalculateProfit calculates the profit and EBT/Profit ratio after tax.
func (investment *Investment) CalculateProfit() (profit, ebtProfitRatio float64) {
	futureRealValue := investment.CalculateFutureRealValue()
	profit = futureRealValue * (1 - investment.TaxRate/100)
	ebtProfitRatio = futureRealValue / profit
	return profit, ebtProfitRatio
}

func main() {
	investment := NewInvestment()

	var err error

	investment.Amount, err = GetFloatInput("Investment amount")
	if err != nil {
		fmt.Println(err)
		return
	}

	investment.ExpectedRate, err = GetFloatInput("Expected return rate (%)")
	if err != nil {
		fmt.Println(err)
		return
	}

	investment.Years, err = GetFloatInput("Number of years")
	if err != nil {
		fmt.Println(err)
		return
	}

	futureValue := investment.CalculateFutureValue()
	futureRealValue := investment.CalculateFutureRealValue()
	profit, profitRatio := investment.CalculateProfit()

	fmt.Println()
	fmt.Println("Future value:", fmt.Sprintf("%.2f", futureValue))
	fmt.Println("Future real value:", fmt.Sprintf("%.2f", futureRealValue))
	fmt.Println()

	fmt.Println("Earnings Before Tax (EBT):", fmt.Sprintf("%.2f", futureRealValue))
	fmt.Println("Profit After Tax:", fmt.Sprintf("%.2f", profit))
	fmt.Println("EBT/Profit Ratio:", fmt.Sprintf("%.2f", profitRatio))
}
