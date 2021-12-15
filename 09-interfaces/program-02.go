package main

import "fmt"

func main() {
	var x interface{}
	//x = "hello"
	x = 10
	//x = true
	if val, ok := x.(int); ok {
		y := val + 100
		fmt.Println(y)
	} else {
		fmt.Println("not an int")
	}

	//type switch
	//x = "hello"
	//x = true
	x = struct{}{}
	switch val := x.(type) {
	case int:
		fmt.Println("x is an int, x + 100 = ", val+100)
	case string:
		fmt.Println("x is a string, len(x) = ", len(val))
	case bool:
		fmt.Println("x is a bool, x = ", val)
	case []int:
		fmt.Println("x is an int slice, len(x) = ", len(val))
	default:
		fmt.Println("unknown type")
	}
}
