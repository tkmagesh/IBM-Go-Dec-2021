package main

import (
	"fmt"
	"time"
)

func main() {
	ch1, ch2 := make(chan string), make(chan string)
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
	for {
		select {
		case s1 := <-ch1:
			fmt.Println(s1)

		case s2 := <-ch2:
			fmt.Println(s2)
		}
	}
}
