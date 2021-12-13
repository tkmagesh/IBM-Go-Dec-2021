package main

import (
	"fmt"
)

func main() {
	var userChoice int
	for {
		userChoice = getUserChoice()
		if userChoice == 5 {
			break
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		processUserChoice(userChoice)
	}
}

func processUserChoice(choice int) {
	var n1, n2, result int
	var operation func(int, int) int
	n1, n2 = getOperands()
	operation = getOperationByChoice(choice)
	result = operation(n1, n2)
	fmt.Printf("Result = %d\n", result)
}

func getOperationByChoice(choice int) func(int, int) int {
	switch choice {
	case 1:
		return add
	case 2:
		return subtract
	case 3:
		return multiply
	case 4:
		return divide
	}
}

func getOperands() (int, int) {
	var n1, n2 int
	fmt.Println("Enter the values")
	fmt.Scanf("%d %d\n", &n1, &n2)
	return n1, n2
}

func getUserChoice() int {
	fmt.Println("Enter the choice :")
	fmt.Println("1. Add")
	fmt.Println("2. Subtract")
	fmt.Println("3. Multiply")
	fmt.Println("4. Divide")
	fmt.Println("5. Exit")

	var userChoice int
	fmt.Scanln(&userChoice)
	return userChoice
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}
