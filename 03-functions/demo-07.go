package main

import "fmt"

func main() {
	increment, decrement := getCounters()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
	fmt.Println(decrement())
}

func getCounters() (func() int, func() int) {
	var counter int
	increment := func() int {
		counter++
		return counter
	}
	decrement := func() int {
		counter--
		return counter
	}
	return increment, decrement

}
