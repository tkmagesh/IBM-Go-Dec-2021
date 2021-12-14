package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	//var nos [5]int
	//var nos [5]int = [5]int{3, 1, 4, 2, 5}
	nos := [5]int{3, 1, 4, 2, 5}
	fmt.Println(nos)

	fmt.Println("Iterating an array [using index]")
	for idx := 0; idx < len(nos); idx++ {
		fmt.Println(idx, nos[idx])
	}

	fmt.Println("Iterating an array [using range]")
	for idx, value := range nos {
		fmt.Println(idx, value)
	}

	fmt.Println("Creating a copy of an array")
	newNos := nos
	nos[0] = 100
	fmt.Println("nos => ", nos)
	fmt.Println("newNos => ", newNos)
	fmt.Printf("Address of nos = %p, newNos = %p\n", &nos, &newNos)

	fmt.Println("")
	fmt.Println("Slice")
	//var products []string
	var products []string = []string{"Pen", "Pencil"}
	fmt.Println(products)
	fmt.Printf("len(products) => %d, cap(products) => %d\n", len(products), cap(products))

	fmt.Println("Adding new items to the slice")
	products = append(products, "Marker")
	fmt.Printf("After adding 1 item, len(products) => %d, cap(products) => %d\n", len(products), cap(products))
	//products = append(products, "Scribble-pad", "mouse-pad")
	newProducts := []string{"Scribble-pad", "mouse-pad"}
	products = append(products, newProducts...)
	fmt.Printf("After adding 2 more items, len(products) => %d, cap(products) => %d\n", len(products), cap(products))
	fmt.Println(products)

	//using make function to proactively allocate memory
	randomNos := make([]int, 0, 100)
	fmt.Println(randomNos, len(randomNos), cap(randomNos))
	randomNos = append(randomNos, 10, 20, 30, 40, 50)
	fmt.Println("After adding 5 random values")
	fmt.Println(randomNos, len(randomNos), cap(randomNos))
}
