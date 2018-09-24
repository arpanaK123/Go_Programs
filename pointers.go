package main
import "fmt"
func main(){
i:=0
fmt.Println("i= ",i)
zeroval(i)
fmt.Println("value of i: ",i)
zeropointer(&i)
fmt.Println("i: ",i)
fmt.Println("pointer of i: ",&i)
}
func zeroval(value int){
value=0
}
func zeropointer(ipointer *int){
	*ipointer=0
}