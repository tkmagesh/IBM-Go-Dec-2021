package main

import (
	"fmt"
	"sync"
)

type OpCount struct {
	count int
	sync.Mutex
}

func (opCount *OpCount) Increment() {
	opCount.Lock()
	{
		opCount.count++
	}
	opCount.Unlock()
}

var opCount OpCount

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1000)
	for idx := 0; idx < 1000; idx++ {
		go add(idx, idx, wg)
	}
	wg.Wait()
	fmt.Println("opCount:", opCount.count)
}

func add(x, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("processing %d and %d\n", x, y)
	opCount.Increment()
	//fmt.Println(x + y)
}
