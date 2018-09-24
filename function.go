package main

import (
	"fmt"
)

var (
	price  = 100
	number = 6
)

func multiply(price, number int) int {
	totalmul := price * number
	return totalmul
}
func calculate(price, number int) int {
	total := price * number
	return total
}

var (
	a = 10
	b = 20
	c = 30
)

func add(a, b, c int) (int, int, int) {
	sum := a + b + c
	mul := a * b * c
	subtract := c - b - a
	return sum, mul, subtract
}

var (
	length = 50
	width  = 50
)

func definedTwoValueButUseOne(length, width int) (int, int) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}
func main() {
	// price := 50
	// number := 6
	total := calculate(50, 6)
	fmt.Println("total: ", total)
	totalmul := multiply(price, number)
	fmt.Println("Multiply: ", totalmul)
	sum, mul, subtract := add(a, b, c)
	fmt.Println("sum: ", sum, "mul: ", mul, "subtract: ", subtract)
	area, _ := definedTwoValueButUseOne(length, width)
	fmt.Println("Area: ", area)
}
