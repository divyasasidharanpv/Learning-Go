package main

import "fmt"

var num int

func main() {
	num, _ = fmt.Println("hello")
	fmt.Println(num)
}

// func Println(a ...any) (n int, err error) >> two return values of type int and error. Need not declare variable for error because we are ignoring it.
// "_" is used to ignore error in this case
