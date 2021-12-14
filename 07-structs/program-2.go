package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

//association
type PerishableProduct struct {
	product Product
	expiry  string
}

func main() {
	pp := PerishableProduct{product: Product{201, "Grapes", 60, 5, "Food"}, expiry: "2 Days"}
	//fmt.Printf("%#v\n", pp)
	fmt.Println(ToString(pp))
}

func ToString(p PerishableProduct) string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %q, Expiry = %s", p.product.Id, p.product.Name, p.product.Cost, p.product.Units, p.product.Category, p.expiry)
}
