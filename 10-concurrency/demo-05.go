package main

import (
	"fmt"
	"sync"
)

var opCount int
var mutex sync.Mutex = sync.Mutex{}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for idx := 0; idx < 1000; idx++ {
		go add(idx, idx, wg)
	}
	wg.Wait()
	fmt.Println("opCount:", opCount)
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("processing %d and %d\n", x, y)
	mutex.Lock()
	{
		opCount++
	}
	mutex.Unlock()
	//fmt.Println(x + y)
}
