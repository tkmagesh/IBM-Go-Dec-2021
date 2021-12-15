package main

import "fmt"

func main() {
	fmt.Println("main started")
	go f1()
	f2()
	fmt.Println("main ended")
}

func f1() {
	fmt.Println("f1 invoked")
}

func f2() {
	fmt.Println("f2 invoked")
}
