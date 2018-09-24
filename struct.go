package main
import "fmt"
type userDetails struct{
	name string
	address string
	age int
}
func main(){
fmt.Println(userDetails{"arya","usa",3})
fmt.Println(&userDetails{"arya","usa",3})

value:=userDetails{name: "arya",address: "usa",age:3}
fmt.Println("name: ",value.name)
fmt.Println("age: ",value.age)
}