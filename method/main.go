package main

import "fmt"

type Item struct {
	name		string
	price		float64
	quantity	int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}
func main() {
	shirt := Item{name: "Men's Slim-Fit shirt", price: 25000, quantity: 3}

	fmt.Println(shirt.Cost())
}