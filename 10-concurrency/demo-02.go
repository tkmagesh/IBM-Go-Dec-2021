package main

import (
	"fmt"
)

func main() {
	fmt.Println("main started")
	go f1()
	f2()

	//DO NOT DO THIS
	/*
		time.Sleep(1 * time.Second)
	*/
	var input string

	//DO NOT DO THIS
	fmt.Scanln(&input)
	fmt.Println("main ended")

}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
