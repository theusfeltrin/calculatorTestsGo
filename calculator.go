package main

import (
	"fmt"
)

func sum(x float64, y float64) float64 {
	return x + y
}
func sub(x float64, y float64) float64 {
	return x - y
}
func mult(x float64, y float64) float64 {
	return x * y
}
func divide(x float64, y float64) float64 {
	return x / y
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Println("Simple Calculator")
	fmt.Println("-----------------")

	fmt.Print("Enter the first number: ")
	_, err1 := fmt.Scanf("%f", &num1)
	if err1 != nil {
		fmt.Println("Invalid input for the first number.")
		return
	}

	fmt.Print("Enter the operator (+, -, *, /): ")
	_, err2 := fmt.Scanf("%s", &operator)
	if err2 != nil {
		fmt.Println("Invalid operator.")
		return
	}

	fmt.Print("Enter the second number: ")
	_, err3 := fmt.Scanf("%f", &num2)
	if err3 != nil {
		fmt.Println("Invalid input for the second number.")
		return
	}

	result := 0.0

	switch operator {
	case "+":
		result = sum(num1, num2)
	case "-":
		result = sub(num1, num2)
	case "*":
		result = mult(num1, num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			return
		}
		result = divide(num1, num2)
	default:
		fmt.Println("Invalid operator.")
		return
	}

	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
