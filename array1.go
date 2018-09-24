package main
import "fmt"
func main(){
 var a[5] int
	fmt.Println("a= ", a)
a[4]=200
fmt.Println("set: ",a)
fmt.Println("get: ",a[4])
fmt.Println("length: ",len(a))

b:=[5]int{1,2,3,4,5}
fmt.Println("b : ",b)
twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
fmt.Println("twod: ",twoD)
}