package main
import "errors"
import "fmt"
func f1(arg int) (int, error) {
    if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	func main(){
		fmt.Println(f1)
	}
	