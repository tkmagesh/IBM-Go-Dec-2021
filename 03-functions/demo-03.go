package main

import "fmt"

func main() {
	/* printHello := func() {
		fmt.Println("Hello world!")
	} */
	var printHello func()
	printHello = func() {
		fmt.Println("hello world!")
	}
	printHello()

	var add func(int, int) int
	add = func(x, y int) int {
		return x + y
	}
	fmt.Println(add(100, 200))

	//anonymous functions
	func() {
		fmt.Println("Printed from an anonymous function")
	}()

	func(s string) {
		fmt.Println("anon : ", s)
	}("Welcome to Golang!")

	func(x, y int) {
		fmt.Println("x + y = ", x+y)
	}(10, 20)

	result := func(x, y int) int {
		return x + y
	}(1000, 2000)
	fmt.Println("result = ", result)
}

/*
	func printHello() {
		fmt.Println("hello world!")
	}
*/

/*
func add(x, y int) int {
	return x + y
}
*/
