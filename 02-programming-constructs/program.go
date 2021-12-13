package main

import "fmt"

func main() {
	//if construct
	/*
		no := 21
		if no%2 == 0 {
			fmt.Println(no, "is even")
		} else {
			fmt.Println(no, "is odd")
		}
	*/

	if no := 21; no%2 == 0 {
		fmt.Println(no, "is even")
	} else {
		fmt.Println(no, "is odd")
	}

	//for construct
	//ver 1.0
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("sum =", sum)

	//ver 2.0
	//sumulate a 'while' construct
	numSum := 1
	for numSum < 100 {
		numSum += numSum
	}
	fmt.Println("numSum =", numSum)

	//ver 3.0
	//infefinite loop
	x := 1
	for {
		if x > 100 {
			break
		}
		x += x
	}
	fmt.Println("x =", x)

	//switch construct
	/*
		rank by score
		score 0 - 3 => "Terrible"
		score 4 - 7 => "Not Bad"
		score 8 - 9 => "Good"
		score 10 => "Excellent"
		otherwise => "Invalid Score"
	*/

	score := 7
	switch score {
	case 0, 1, 2, 3:
		fmt.Println("Terrible")
	case 4, 5, 6, 7:
		fmt.Println("Not Bad")
	case 8, 9:
		fmt.Println("Good")
	case 10:
		fmt.Println("Excellent")
	default:
		fmt.Println("Invalid scorez!")
	}

	newNo := 22
	/*
		switch newNo % 2 {
		case 0:
			fmt.Println(newNo, "is even")
		case 1:
			fmt.Println(newNo, "is odd")
		}
	*/

	switch {
	case newNo%2 == 0:
		fmt.Println(newNo, "is even")
	case newNo%2 == 1:
		fmt.Println(newNo, "is odd")
	}

	//falthrough
	n := 4
	switch n {
	case 0:
		fmt.Println("is zero")
	case 1:
		fmt.Println("n <= 1")
		fallthrough
	case 2:
		fmt.Println("n <= 2")
		fallthrough
	case 3:
		fmt.Println("n <= 3")
		fallthrough
	case 4:
		fmt.Println("n <= 4")
		fallthrough
	case 5:
		fmt.Println("n <= 5")
		fallthrough
	case 6:
		fmt.Println("n <= 6")
	case 7:
		fmt.Println("n <= 7")
		fallthrough
	case 8:
		fmt.Println("n <= 8")
	}
}
