package main

import "fmt"

func find(num int, nums ...int) {
	fmt.Println("type of nums is ", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, " number found a index", i, "in", num)
			found = true
		}

	}

	if !found {
		fmt.Println(num, " num not found ", nums)
	}
	fmt.Printf("\n")

}

func change(s ...string) {
	s[0] = "mumbai"
	s[2] = "maharastra"
}
func appendString(str ...string) {
	str[0] = "India"
	str = append(str, "Bangalore")
	fmt.Println("string after append: ", str)
}
func main() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87, 80)
	find(10)

	// find a slice to a variadic tion
	nums := []int{23, 46, 79}
	find(46, nums...)
	// nums := []int{89, 90, 95}
	// find(90, nums)//error

	welcome := []string{"hello", "Bridgelabz", ""}
	change(welcome...)
	fmt.Println("welcome---> ", welcome)
	fmt.Println("---------------------------------")
	welcome2 := []string{"hello", "---", "---"}
	appendString(welcome...)
	fmt.Println("after: ", welcome2)
}
