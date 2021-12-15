package main

import (
	"fmt"
	"time"
)

func main() {

	resultCh := make(chan int)
	go add(10, 20, resultCh)
	result := <-resultCh
	fmt.Println("result:", result)
}

func add(x, y int, ch chan int) {
	time.Sleep(3 * time.Second)
	fmt.Printf("processing %d and %d\n", x, y)
	result := x + y
	ch <- result
}
