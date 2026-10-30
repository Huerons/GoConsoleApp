package main

import (
	"fmt"
	"strconv"
)

func main() {

calculator:
	for {
		fmt.Println("Hello, World! This is a simple Go application.")
		fmt.Println("Calculator")
		fmt.Println("1. Addition")
		fmt.Println("2. Subtraction")
		fmt.Println("3. Multiplication")
		fmt.Println("4. Division")
		var choice int
		fmt.Print("Enter your choice (1-4): ")
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			var a, b float64
			fmt.Print("Enter the first number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&a)
			aRes, aErr := strconv.ParseFloat(strconv.FormatFloat(a, 'f', -1, 64), 64)
			fmt.Print("Enter the second number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&b)
			bRes, bErr := strconv.ParseFloat(strconv.FormatFloat(b, 'f', -1, 64), 64)
			if aErr != nil || bErr != nil {
				fmt.Println("Error converting numbers")
			} else {
				fmt.Printf("Result: %g\n", aRes+bRes)
			}
		case 2:
			var a, b float64
			fmt.Print("Enter the first number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&a)
			fmt.Print("Enter the second number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&b)
			aRes, aErr := strconv.ParseFloat(strconv.FormatFloat(a, 'f', -1, 64), 64)
			bRes, bErr := strconv.ParseFloat(strconv.FormatFloat(b, 'f', -1, 64), 64)
			if aErr != nil || bErr != nil {
				fmt.Println("Error converting numbers")
			} else {
				fmt.Printf("Result: %g\n", aRes-bRes)
			}
		case 3:
			var a, b float64
			fmt.Print("Enter the first number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&a)
			fmt.Print("Enter the second number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&b)
			aRes, aErr := strconv.ParseFloat(strconv.FormatFloat(a, 'f', -1, 64), 64)
			bRes, bErr := strconv.ParseFloat(strconv.FormatFloat(b, 'f', -1, 64), 64)
			if aErr != nil || bErr != nil {
				fmt.Println("Error converting numbers")
			} else {
				fmt.Printf("Result: %g\n", aRes*bRes)
			}
		case 4:
			var a, b float64
			fmt.Print("Enter the first number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&a)
			fmt.Print("Enter the second number (Enter a valid number [No Value is equal to zero]): ")
			fmt.Scanln(&b)
			aRes, aErr := strconv.ParseFloat(strconv.FormatFloat(a, 'f', -1, 64), 64)
			bRes, bErr := strconv.ParseFloat(strconv.FormatFloat(b, 'f', -1, 64), 64)
			if aErr != nil || bErr != nil {
				fmt.Println("Error converting numbers")
			} else {
				fmt.Printf("Result: %g\n", aRes/bRes)
			}
		default:
			fmt.Println("Invalid choice.")
			break calculator
		}
	}
}
