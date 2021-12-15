package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go print("Hello", ch1, ch2, wg)
	go print("World", ch2, ch1, wg)
	ch1 <- "start"
	wg.Wait()
}

func print(s string, in, out chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		<-in
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
		out <- "done"
	}
	wg.Done()
}

/*
// Expected Output:
// Hello
// World
// Hello
// World
// Hello
// World
// Hello
// World
// Hello
// World
*/
