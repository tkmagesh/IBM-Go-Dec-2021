package main

import "fmt"

func main() {
	var i int
	fmt.Println("Enter the no :")
	fmt.Scanln(&i)

	isPrime := true
	for j := 2; j <= (i / 2); j++ {
		if i%j == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		//fmt.Println(i, "is prime")
		fmt.Printf("%d is prime\n", i)
	} else {
		fmt.Printf("%d is not prime\n", i)
	}

}
