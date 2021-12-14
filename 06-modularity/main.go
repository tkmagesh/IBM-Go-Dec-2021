package main

import (
	"fmt"
	calculator "modularity-demo/calc"

	"modularity-demo/calc/utils"

	"github.com/fatih/color"
)

func main() {
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	fmt.Println("OpCount = ", calculator.GetOpCount())
	fmt.Println("Is 21 odd ? : ", utils.IsOdd(21))
	color.Red("This will be printed in red")
}
