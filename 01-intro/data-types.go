package main

import "fmt"

var x int = 100

func main() {
	/*
		var msg string
		msg = "Hello World!"
	*/

	/*
		var msg string = "Hello World!"
	*/

	/* Type inference */
	/*
		var msg = "Hello World!"
	*/

	msg := "Hello World!" /* Note => ":=" syntax is applicable only in a function */
	fmt.Println(msg)
}
