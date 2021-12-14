package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float64
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	expiry string
}

func main() {
	pen := &Product{102, "Pen", 10, 100, "Stationary"}
	//fmt.Println(ToString(pen))
	fmt.Println(pen.ToString())
	//applyDiscount(&pen, 10)
	pen.applyDiscount(10)
	fmt.Println(pen.ToString())

	pp := PerishableProduct{Product{200, "Grapes", 60, 15, "Fruit"}, "2 Days"}
	fmt.Println("Perishable Product")
	fmt.Println(pp.ToString())
}

func (p Product) ToString() string {
	return fmt.Sprintf("Id = %d, Name = %s, Cost = %v, Units = %d, Category = %q", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

//Write a function (applyDiscount) that can used to apply discount(%) for a product
func (p *Product) applyDiscount(discount float64) {
	p.Cost = p.Cost * ((100 - discount) / 100)
}

func (pp PerishableProduct) ToString() string {
	return fmt.Sprintf("%s, expiry = %s", pp.Product.ToString(), pp.expiry)
}
