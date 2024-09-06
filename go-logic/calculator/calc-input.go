package main

import (
	"fmt"
)

func main() {
	var operand1, operand2 float64
	var operator string

	fmt.Println("Calculator")
	fmt.Print("Enter first number: ")
	fmt.Scan(&operand1)

	fmt.Print("Enter operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter second number: ")
	fmt.Scan(&operand2)

	var result float64
	switch operator {
	case "+":
		result = operand1 + operand2
	case "-":
		result = operand1 - operand2
	case "*":
		result = operand1 * operand2
	case "/":
		if operand2 == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		result = operand1 / operand2
	default:
		fmt.Println("Invalid operator. Available operators: +, -, *, /")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
