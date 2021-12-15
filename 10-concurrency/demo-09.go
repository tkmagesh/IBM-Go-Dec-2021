package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	resultCh := make(chan int)
	go add(10, 20, wg, resultCh)
	result := <-resultCh
	wg.Wait()
	fmt.Println("result:", result)
}

func add(x, y int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	time.Sleep(3 * time.Second)
	fmt.Printf("processing %d and %d\n", x, y)
	result := x + y
	ch <- result
	//fmt.Println(x + y)
}
