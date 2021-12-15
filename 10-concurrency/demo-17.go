package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go fibonacci(ch)
	for no := range ch {
		fmt.Println(no)
	}
	fmt.Println("Exiting main")
}

func fibonacci(ch chan int) {
	x, y := 0, 1
	done := func() chan bool {
		doneCh := make(chan bool)
		go func() {
			time.Sleep(time.Second * 15)
			doneCh <- true
		}()
		return doneCh
	}()
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
