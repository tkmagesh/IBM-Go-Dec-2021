package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		fmt.Println("Enter the choice :")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		var userChoice int
		fmt.Scanln(&userChoice)
		if userChoice == 5 {
			os.Exit(0)
		}
		if userChoice < 1 || userChoice > 5 {
			fmt.Println("Invalid choice")
			continue
		}
		var n1, n2, result int
		fmt.Println("Enter the values")
		fmt.Scanf("%d %d\n", &n1, &n2)
		switch userChoice {
		case 1:
			result = n1 + n2
		case 2:
			result = n1 - n2
		case 3:
			result = n1 * n2
		case 4:
			result = n1 / n2
		}
		fmt.Printf("Result = %d\n", result)

	}
}
