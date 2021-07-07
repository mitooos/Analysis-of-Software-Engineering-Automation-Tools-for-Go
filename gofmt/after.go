package gofmt

import "fmt"

var a int = 2
var b int = 5
var c string = `hello world`
var list = []int{
	1, 2, 3,
	4,
	5,
}

func print() {
	fmt.Println("Value for a,b and c is : ")
	fmt.Println(a)
	fmt.Println((b))
	fmt.Println(c)
}
