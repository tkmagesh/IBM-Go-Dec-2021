package main

import "fmt"

func main() {
	data := []int{4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55, 4, 1, 7, 3, 8, 5, 9, 3, 11, 15, 63, 87, 29, 40, 50, 55}

	/*
		split the data into 2 equal sets
		aggregate the data sets using the sum function (should happen concurrently)
		collect the intermediate result
		aggregate the intermediate result
		print the final result
	*/

	dataSet1 := data[:len(data)/2]
	dataSet2 := data[len(data)/2:]
	ch := make(chan int)
	go sum(ch, dataSet1...)
	go sum(ch, dataSet2...)
	result := <-ch + <-ch
	fmt.Println("Result : ", result)
}

func sum(resultCh chan int, nos ...int) {
	var result int
	for _, no := range nos {
		result += no
	}
	resultCh <- result
}
