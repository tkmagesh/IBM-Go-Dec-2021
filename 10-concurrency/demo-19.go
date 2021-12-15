package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go fibonacci(ch, done)
	go func() {
		var input string
		fmt.Scanln(&input)
		done <- true
	}()
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Exiting main")
}

func fibonacci(ch chan int, done chan bool) {
	x, y := 0, 1
	for {
		select {
		case ch <- x:
			time.Sleep(500 * time.Millisecond)
			x, y = y, x+y
		case <-done:
			fmt.Println("Done")
			close(ch)
			return
		}
	}
}
