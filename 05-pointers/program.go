package main

import "fmt"

func main() {
	var no int = 100
	var noPtr *int = &no
	fmt.Println(noPtr, no)
	fmt.Println("Deferencing")
	newNo := *noPtr
	fmt.Println(newNo)

	fmt.Println()
	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 10, 20
	fmt.Println("before swapping ", x, y)
	swap(&x, &y)
	fmt.Println("after swapping ", x, y)

}

func increment(val *int) {
	*val++
}

func swap(no1Ptr, no2Ptr *int) {
	*no1Ptr, *no2Ptr = *no2Ptr, *no1Ptr
}
