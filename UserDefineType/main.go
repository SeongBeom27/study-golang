package main

import (
	"fmt"
)

// -------------------------------- 1. 구조체 ------------------------------
type Item struct {
	name		string
	price		float64
	quantity	int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

// -------------------------------- 2. 함수 서명을 사용자 정의 타입으로 사용 ------------------------------
type quantity int
type costCalculator func(quantity, float64) float64

func describe(q quantity, price float64, c costCalculator) {
	fmt.Printf("quantity: %d, price: %0.0f, cost:%0.0f\n", 
		q, price, c(q, price))
}

// -------------------------------- 3. 구조체 ------------------------------
type rect struct {
	width	float64
	height	float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

// -------------------------------- 4. 인터페이스 ------------------------------
type shaper interface {
	area()	float64
}

func describe2(s shaper) {
	fmt.Println("area :", s.area())
}

func main() {
// -------------------------------- 1. 구조체 ------------------------------
	shirt := Item{name: "Men's Slim-Fit shirt", price: 25000, quantity: 3}

	fmt.Println(shirt.Cost())

// -------------------------------- 2. 함수 서명을 사용자 정의 타입으로 사용 ------------------------------
	var offBy10Percent, offBy1000Won costCalculator

	offBy10Percent = func(q quantity, price float64) float64 {
		return float64(q) * price * 0.9
	}

	offBy1000Won = func(q quantity, price float64) float64 {
		return float64(q) * price - 1000
	}

	describe(3, 10000, offBy10Percent)
	describe(3, 10000, offBy1000Won)

// -------------------------------- 3. 구조체 ------------------------------
	r := rect{3, 4}
	fmt.Println("area : ", r.area())

// -------------------------------- 4. 인터페이스 ------------------------------
	r2 := rect{3, 4}
	// shaper 인터페이스와 rect 타입 사이에는 아무런 연결 고리가 없음에도 rect 타입이 shaper 인터페이스에 정의된 메서드 (area() float64)와 형태가 같은 메서드를 가진 것만으로도 rect 타입을 shaper 인터페이스로 사용 가능
	describe2(r2)
}