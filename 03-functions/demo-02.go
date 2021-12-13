package main

import "fmt"

func main() {
	fmt.Println(sum("abc", 10))
	fmt.Println(sum("abc", 10, 20))
	fmt.Println(sum("abc", 10, 20, 30))
	fmt.Println(sum("abc", 10, 20, 30, 40))
}

func sum(s string, nos ...int) int {
	//nos => list of int
	//len(nos) => # of values in the list
	//nos[0] => first value in the list
	result := 0
	for i := 0; i < len(nos); i++ {
		result += nos[i]
	}
	return result
}
