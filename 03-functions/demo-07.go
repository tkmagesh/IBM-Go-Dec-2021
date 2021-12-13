package main

import "fmt"

var counter int

func main() {
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	doSomething()
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func increment() int {
	counter++
	return counter
}

func doSomething() {
	counter = 1000
}
