package main

import "fmt"

func main() {
	fmt.Printf("is 59 prime ? %t\n", isPrime(59))
	fmt.Printf("is 97 prime ? %t\n", isPrime(97))
	fmt.Printf("is 95 prime ? %t\n", isPrime(95))

	fmt.Printf("add(100,200) = %d\n", add(100, 200))
	//fmt.Println(divide(100, 7))
	/*
		quotient, remainder := divide(100, 7)
		fmt.Println(quotient, remainder)
	*/

	quotient, _ := divide(100, 7)
	fmt.Println(quotient)
}

/*
func add(x int, y int) int {
	return x + y
}
*/

/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

func add(x, y int) (result int) {
	result = x + y
	return
}

func divide(x, y int) (quotient int, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}

func isPrime(n int) bool {
	for i := 2; i <= (n / 2); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
