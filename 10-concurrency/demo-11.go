package main

import (
	"fmt"
	"time"
)

func main() {
	evenNosCh := genEvenNos()
	for idx := 0; idx < 20; idx++ {
		evenNo := <-evenNosCh
		fmt.Println(evenNo)
	}
}

func genEvenNos() <-chan int {

	resultCh := make(chan int)
	go func() {
		var no int
		for {
			time.Sleep(500 * time.Millisecond)
			no += 2
			fmt.Println("Generated:", no)
			resultCh <- no
		}
	}()

	return resultCh
}
