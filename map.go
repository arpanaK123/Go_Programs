package main

import "fmt"

func main() {
	var personSalary map[string]int
	fmt.Println("map is nil", personSalary)
	personSalary = make(map[string]int)
	fmt.Println("person salary: ", personSalary)
	personSalary["bridgelabz"] = 1000
	personSalary["mumbai"] = 100000000000
	personSalary["India"] = 12345678900
	fmt.Println("perons: ", personSalary)
}
