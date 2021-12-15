package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	fmt.Println("main started")
	go f1()
	f2()
	wg.Wait() //blocked until the wg counter becomes 0
	fmt.Println("main ended")

}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(5 * time.Second)
	fmt.Println("f1 ended")
	wg.Done()
}

func f2() {
	fmt.Println("f2 invoked")
}
