package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	b := [5]int{10, 20, 30, 40, 50}
	fmt.Println(b)
	c := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(c)

	d := [3]int{1, 2, 3}
	var e [3]int
	e = d
	fmt.Println(e)
	name := [...]string{"a", "b", "o", "c", "d", "e", "f", "g", "h"}
	fmt.Println(name)
	copy := name
	fmt.Println("copy: ", copy)

	value := [...]int{3, 2, 4, 5, 1, 5, 7, 9}
	for i := 0; i < len(value); i++ {
		fmt.Printf("%d th element of a is %.2f\n", i, value[i])
	}
	fmt.Println("-----------------------------------")
	for i, v := range value { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
	}
	fmt.Println("_________________________________________")
	for _, v := range value {
		fmt.Println("value: ", v)
	}
	array := [2][2]string{
		{"lion", "tiger"},
		{"dog", "cat"},
	}
	arrayPrrint(array)

	var multiarray [2][2]string
	multiarray[0][0] = "cat"
	multiarray[0][1] = "dog"
	multiarray[1][0] = "cat"
	multiarray[1][1] = "dog"

	fmt.Println(multiarray)

	//slice
	slice := [9]int{0, 8, 7, 6, 4, 5, 9, 3, 1}
	var slice1 []int = slice[2:5] //slice between slice[1] to slice[3] ie,[8,7,6]
	fmt.Println("slice: ", slice1)
	fmt.Println(slice)

	fmt.Println("------------------------------")
	fmt.Println("slice before--> ", slice)
	for i := range slice {
		slice[i]++
	}
	fmt.Println("--slice---> ", slice)

	num := [3]int{80, 85, 90}
	num1 := num[:] //creates a slice which contains all elements of the array
	num2 := num[:]
	fmt.Println("array before change: ", num)
	num1[0] = 100
	fmt.Println("array after modification: ", num)
	num2[1] = 101
	fmt.Println("array after:  -->", num)

	//length and capacity of slice
	fruitarray := [...]string{"apple", "banana", "orange", "Pineapple", "litchi", "coconut"}
	fruit := fruitarray[1:4]
	fmt.Println("fruitarray before: ", fruitarray)
	fmt.Println("after--------->")
	fmt.Println("----->fruit", fruit)
	fmt.Println("fruit length--->", len(fruit), "capacity--->", cap(fruit))

	//creation a slice using make func make([]T, len, cap)
	makeslice := make([]int, 5, 5)
	fmt.Println("makeslice--> ", makeslice)

	//using append function
	car := []string{"nano", "safari", "bmw", "honda"}
	fmt.Println("car: ", car, "car len: ", len(car), "car cap: ", cap(car))
	car = append(car, "farrari")
	fmt.Println("after append function: ", car, "car len: ", len(car), "car cap: ", cap(car))

	var name1 []string
	if name1 == nil {
		fmt.Println("name ", name1)
		name1 = append(name1, "asdf", "lkjhg")
		fmt.Println("name1---> ", name1)
	}

	//append one slice to another using the ... operator.
	city := []string{"bangalor", "mumbai", "Delhi"}
	fruits := []string{"Apple", "orange", "bnana"}
	appendsBothString := append(city, fruits...)
	fmt.Println("append both the string: --->", appendsBothString)

	numbers := []int{10, 20, 30}
	fmt.Println("numbers before:---> ", numbers)
	subtractOne(numbers)
	fmt.Println("number after----> ", numbers)

	language := [][]string{
		{"go ", " java "},
		{"javascript"},
		{"html ", " css"},
		{"aws "},
	}
	//multidimensional slice

	for _, language1 := range language {
		for _, language2 := range language1 {
			fmt.Printf("%s ", language2)
		}
		fmt.Printf("\n")
	}

}

//multidimensional array
func arrayPrrint(array [2][2]string) {

	for _, v := range array {
		for _, v1 := range v {
			fmt.Printf(" %s ", v1)
		}
		fmt.Printf("\n")
	}
}

func subtractOne(number []int) {
	for j := range number {
		number[j] -= 2
	}
}
