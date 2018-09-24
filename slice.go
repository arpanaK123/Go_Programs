package main
import "fmt"
func main(){
	a:=make([]string, 3)
	fmt.Println("a: ",a)
	a[0]="1"
	a[1]="bridgelabz"
	a[2]="mumbai"
	fmt.Println("a value: ",a)
fmt.Println("get: ",a[1])
a=append(a,"3")
a=append(a,"4")
fmt.Println(a)

c:=make([]string, len(a))
copy(c,a)
fmt.Println("copy: ",c)
//Slices support a “slice” operator with the syntax slice[low:high]
slice:=a[0:5]
fmt.Println("slice: ",slice)
slice =a[:5]
fmt.Println("slice new: ",slice)
slice =a[2:]
fmt.Println("slice new 1: ",slice)

str:=[]string{"a","s","d"}
fmt.Println("str: ",str)

twoD:=make([][]int,3)
for i := 0; i < 3; i++ {
	innerLen := i + 1
	twoD[i] = make([]int, innerLen)
	for j := 0; j < innerLen; j++ {
		twoD[i][j] = i + j
	}
	fmt.Println("innerLen:",innerLen )
}
fmt.Println("add: ", twoD)
}