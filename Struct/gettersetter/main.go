package main

import "fmt"

type Item struct {
	name		string		
	quantity	int			
	price		float64		
}

func (t *Item) Name() string {
	return t.name
}

func (t *Item) SetName(n string) {
	if len(n) != 0 {
		t.name = n
	}
}

func (t *Item) Quantity() int {
	return t.quantity
}

func (t *Item) SetQuantity(q int) {
	t.quantity = q
}

func (t *Item) Price() float64 {
	return t.price
}

func (t *Item) SetPrice(p float64) {
	if p > 0 {
		t.price = p
	}
}


func main() {
	shirt := Item{"marvel", 1, 2500}

	shirt.SetName("iron")
	shirt.SetQuantity(3)
	shirt.SetPrice(3000)

	fmt.Println(shirt.Name())
	fmt.Println(shirt.Quantity())
	fmt.Println(shirt.Price())
}