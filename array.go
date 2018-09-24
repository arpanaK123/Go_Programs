package main 

import (  
    "fmt"
)

// func main() {  
// 	var b [3]int 
// fmt.Println("b ",b)
// 	var a [3]int //int array with length 3
//     a[0] = 12 
//     a[1] = 78
//     a[2] = 50

//     //a := [3]int{12, 78, 50} 
// 	fmt.Println(a)
// }
func main() {  
    a := [...]string{"USA", "China", "India", "Germany", "France"}
    b := a // a copy of a is assigned to b
    b[1] = "Singapore"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}