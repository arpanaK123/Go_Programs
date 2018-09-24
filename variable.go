package main

import "fmt"

func main() {
	var firstname, lastname, address = "ariya", "reha", "newyork"
	fmt.Println(firstname, lastname, "from", address)
	var name, age = "monika", 23
	fmt.Println("name", name, "age", age)

	var (
		Name    = "Riya"
		Age     = 24
		Address = "Bngalore"
	)
	fmt.Println("Name: ", Name, "Age:", Age, "Address: ", Address)
	var (
		a = 10
		b = 20
		c = 30
	)
	fmt.Println("sum: ", a+b+c)
	var (
		length = 20.20
		bredth = 20.20
		height = 20.20
	)
	fmt.Println("ToTal: ", length+bredth+height)
	var (
		value1 = 20.0
		value2 = 20.0
		value3 = 50.50
	)
	fmt.Println("Value:= ", value1+value2+value3)

	var (
		a1 = 20
		a2 = 30.20
	)
	fmt.Println("Sum: ", a1+int(a2))
	fmt.Println("Sum:", float64(a1)+a2)

}
