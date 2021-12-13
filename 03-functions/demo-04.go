package main

import "fmt"

func main() {
	printHello := fn()
	printHello()
}

func fn() func() {
	return func() {
		fmt.Println("Hello World!")
	}
}
