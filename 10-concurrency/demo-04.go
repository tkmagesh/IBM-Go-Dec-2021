package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	wg.Add(1)
	fmt.Println("main started")
	go f1(wg)
	f2()
	wg.Wait() //blocked until the wg counter becomes 0
	fmt.Println("main ended")

}

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 ended")

}

func f2() {
	fmt.Println("f2 invoked")
}
