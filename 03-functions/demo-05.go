package main

import "fmt"

func main() {
	processOp(add, 100, 200)
	processOp(subtract, 100, 200)
}

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func processOp(oper func(int, int) int, n1 int, n2 int) {
	fmt.Println("Before invocation")
	result := oper(n1, n2)
	fmt.Println("Result = ", result)
	fmt.Println("After invocation")
}
