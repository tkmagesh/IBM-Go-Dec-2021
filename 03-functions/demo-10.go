package main

import (
	"fmt"
	"log"
)

func main() {
	defer func() {
		fmt.Println("Deffered from main!")
		if err := recover(); err != nil {
			fmt.Println("recovering from a panic")
			log.Println(err)
		}
	}()
	fmt.Println(divide(100, 0))
}

func divide(x, y int) (int, int) {
	defer func() {
		fmt.Println("Deferred from divide")
		if err := recover(); err != nil {
			fmt.Println("recovering from a divide")
			log.Println(err)
		}
	}()
	return x / y, x % y
}
