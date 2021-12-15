package main

import (
	"fmt"
	"sync"
	"time"
)

//Not advisable to use global variables
var result int

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go add(10, 20, wg)
	wg.Wait()
	fmt.Println("result:", result)
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Printf("processing %d and %d\n", x, y)
	result = x + y
	//fmt.Println(x + y)
}
