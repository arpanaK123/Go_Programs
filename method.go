package main
import "fmt"
type add struct{
	a int
	b int
}
func(addition *add) addnumber()int{
	return addition.a+addition.b
}
func(mul add) mulnumber()int{
	return mul.a*mul.b
}
func main(){
totalsum:=add{a:10,b:10}
fmt.Println("totalsum",totalsum.addnumber())
fmt.Println("totalmul",totalsum.mulnumber())

}