package main

import "fmt"

var a, b, c int8

func main() {
	a = 10
	b = 20
	c = 30
	fmt.Println("Values of a, b and c are: ", a, b, c)
}

func variable() {
	a = 1
	b = 2
	c = 3
	fmt.Println("Values of a, b and c are: ", a, b, c)
}
