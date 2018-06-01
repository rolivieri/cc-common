package main

//$ echo $GOPATH
// Users/olivieri/go:/Users/olivieri/git/go-inner-folder
//$ go build
import (
	"fmt"
	"myproject/mylib"
)

func main() {
	fmt.Println("Hello World!")
	user := mylib.User{Firstname: "John", Lastname: "Doe", Email: "john@us.ibm.com"}
	fmt.Println(user)
}
