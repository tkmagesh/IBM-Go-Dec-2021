package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

func main() {
	//var p Product
	//p := Product{}
	//var p Product = Product{101, "Pen", 10, 100, "Stationary"}
	//var p Product = Product{101, "Pen", 10}
	//var p Product = Product{Id: 100, Name: "Pen", Cost: 10}
	var p Product = Product{
		Cost: 10,
		Id:   100,
		Name: "Pen",
	}
	//fmt.Printf("%#v\n", p)
	fmt.Println(ToString(p))

	var productPtr *Product = &p

	//deferencing to access the attribute of a struct
	fmt.Println((*productPtr).Id) //=> 100

	//access the attribute directly using the pointer reference
	fmt.Println(productPtr.Id) //=> 100

	applyDiscount(&p, 10)
	fmt.Println(ToString(p)) //=> Cost 9
}

func ToString(p Product) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

//Write a function (applyDiscount) that can used to apply discount(%) for a product
func applyDiscount(p *Product, discount float64) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}
