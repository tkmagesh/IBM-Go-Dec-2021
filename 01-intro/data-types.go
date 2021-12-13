package main

import (
	"fmt"
)

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

	//fmt.Println(strconv.IntSize)

	//Handling multiple variables
	/*
		var x int
		var y int
		var result int
		var message string

		x = 100
		y = 200
		result = x + y
		message = "Result ="
	*/

	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var message string = "Result ="
	*/

	/*
		var x, y, result int
		var message string
		x = 100
		y = 200
		result = x + y
		message = "Result ="
	*/
	/*
		var (
			x, y, result int
			message      string
		)
		x = 100
		y = 200
		result = x + y
		message = "Result ="
	*/
	/*
		var (
			x, y, result int    = 100, 200, 0
			message      string = "Result ="
		)
		result = x + y
	*/

	/*
		x, y := 100, 200
		message := "Result ="
		result := x + y
	*/
	x, y, message := 100, 200, "Result ="
	result := x + y
	fmt.Println(message, result)
}
