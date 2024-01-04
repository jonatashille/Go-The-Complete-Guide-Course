package main

import (
	"errors"
	"fmt"
	"os"
)

// Goals
// 1) Validate user input
//		=> Show error message & exit if invalid input is provided
//		- No negative numbers
//		- Not 0
//	2) Store calculated results into file

const calculatorFile = "calcFile.txt"

func writeFile(ebt, profit, ratio float64) {
	fileText := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f", ebt, profit, ratio)
	os.WriteFile(calculatorFile, []byte(fileText), 0644)
}

func main() {

	revenue, err1 := getUserInput("Revenue: ")
	expenses, err2 := getUserInput("Expenses: ")
	taxRate, err3 := getUserInput("Tax Rate: ")

	if err1 != nil || err2 != nil || err3 != nil {
		fmt.Println(err1)
		panic(err1)
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	writeFile(ebt, profit, ratio)
	fmt.Println()
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt - (revenue * taxRate / 100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("The value has to be greater than 0.")
	}
	return userInput, nil
}
