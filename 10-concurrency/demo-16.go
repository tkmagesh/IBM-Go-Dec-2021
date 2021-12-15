package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	done := make(chan bool)
	go func() {
		for {
			ch1 <- "from 1"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			ch2 <- "from 2"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		var input string
		fmt.Scanln(&input)
		done <- true
	}()
	stop := false
	for !stop {
		select {
		case s1 := <-ch1:
			fmt.Println(s1)

		case s2 := <-ch2:
			fmt.Println(s2)

		case <-done:
			fmt.Println("Done")
			stop = true
			break
		}
	}
	fmt.Println("Exiting main")
}
