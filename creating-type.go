package main

import "fmt"

var a int

type number int

var b number

func main() {
	a = 10
	b = 20
	fmt.Printf("%d\t", a)
	fmt.Printf("%d,%T\t\n", b, b)
}

// A type "number" is created ; b is of type number
