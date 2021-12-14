package main

import "fmt"

func main() {
	var no int = 100
	var noPtr *int = &no
	fmt.Println(noPtr, no)
	fmt.Println("Deferencing")
	newNo := *noPtr
	fmt.Println(newNo)

	fmt.Println()
	fmt.Println("Before incrementing, no = ", no)
	increment(no)
	fmt.Println("After incrementing, no = ", no)
}

func increment(val int) {
	val++
}
