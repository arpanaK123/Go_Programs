package main
import "fmt"
type number interface{
	add() int
}
type numberstructure struct{
	a ,b int
}
func  (num numberstructure)add()int{
	return num.a+num.b
}
func addition(n number)  {
	fmt.Println("number: ",n)
	fmt.Println("add: ",n.add())
}
func main(){
	num:=numberstructure{a:10,b:10}
	addition(num)
	//fmt.Println(num)
}