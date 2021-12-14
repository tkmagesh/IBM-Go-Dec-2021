package main

import (
	"errors"
	"fmt"
	"log"
)

var DivideByZero = errors.New("division by zero")

func main() {
	defer func() {
		fmt.Println("Deffered from main!")
		if err := recover(); err != nil {
			fmt.Println("recovering from a panic")
			log.Println(err)
		}
	}()
	q, r, err := divideClient(100, 0)
	if err != nil {
		fmt.Println("Error occurred. Taking an alternative path")
		fmt.Println("Do something....")
		return
	}
	fmt.Println(q, r)
}

func divideClient(x, y int) (quotient int, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Recovering from panic : ", e)
			err = DivideByZero
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

func divide(x, y int) (int, int) {
	defer func() {
		fmt.Println("Deferred from divide")
	}()
	if y == 0 {
		panic("cannot divide by zero")
	}
	return x / y, x % y
}
