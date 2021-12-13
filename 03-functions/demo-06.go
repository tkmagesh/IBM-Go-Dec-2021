package main

import "fmt"

func main() {
	/*
		add := getLoggedOp(add)
		subtract := getLoggedOp(subtract)
		add(100, 200)
		subtract(100, 200)
	*/
	getLoggedOp(add)(100, 200)
	getLoggedOp(subtract)(100, 200)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func getLoggedOp(oper func(int, int) int) func(int, int) {
	return func(n1, n2 int) {
		fmt.Println("Before invocation")
		result := oper(n1, n2)
		fmt.Println("Result = ", result)
		fmt.Println("After invocation")
	}
}
