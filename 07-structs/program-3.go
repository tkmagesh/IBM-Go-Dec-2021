package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//composition
type PerishableProduct struct {
	Product
	expiry string
}

func main() {
	pp := PerishableProduct{Product: Product{201, "Grapes", 60, 5, "Food"}, expiry: "2 Days"}
	//fmt.Printf("%#v\n", pp)
	fmt.Println(ToString(pp))
}

/* func ToString(p PerishableProduct) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %q, Expiry = %s", p.Product.Id, p.Product.Name, p.Product.Cost, p.Product.Units, p.Product.Category, p.expiry)
} */

func ToString(p PerishableProduct) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %q, Expiry = %s", p.Id, p.Name, p.Cost, p.Units, p.Category, p.expiry)
}
