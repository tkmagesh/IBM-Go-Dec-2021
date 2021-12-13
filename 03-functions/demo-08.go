package main

import (
	"errors"
	"fmt"
)

func main() {
	q, r, err := divide(100, 7)
	if err != nil {
		fmt.Println("something went wrong", err)
		return
	}
	fmt.Println(q, r)
}

func divide(x, y int) (quotient int, remainder int, err error) {
	if y == 0 {
		err = errors.New("invalid arguments. y cannot be zero")
		return
	}
	quotient, remainder = x/y, x%y
	return
}
