package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func (product Product) Format() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %f, Units = %d, Category = %s", product.Id, product.Name, product.Cost, product.Units, product.Category)
}

type Products []Product

func (products Products) Format() string {
	var result string
	for _, p := range products {
		result += fmt.Sprintf("%v\n", p.Format())
	}
	return result
}

func main() {

	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	fmt.Println(products.Format())

	//implement the sorting for the products.
	//sort the products by id / name / cost / units / category
	//IMPORTANT : sort.Sort() function

}
