package main

import (
	"fmt"
	"time"
)

func main() {
	print("Hello")
	print("World")
}

func print(s string /* other arguments */) {
	for i := 0; i < 5; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(s)
	}
}

/*
// Expected Output:
// Hello
// World
// Hello
// World
// Hello
// World
// Hello
// World
// Hello
// World
*/
