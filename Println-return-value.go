package main

import "fmt"

var num int
var err error

func main() {
	num, err = fmt.Println("Hello")
	fmt.Println(num, err)
}

//
// func Println(a ...any) (n int, err error) >> two return values of type int and error ; so use two variables
