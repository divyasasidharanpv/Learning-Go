package main

import "fmt"

var (
	b float32 = 1.1234
	a int8    = 10
	c         = "String"
	d         = 5
)

func main() {
	fmt.Printf("values of a, b  c and d are %d  %f  %q and %b\n", a, b, c, d)
}

//  %q	a double-quoted string safely escaped with Go syntax
// %f	decimal point but no exponent, e.g. 123.456
// %d	base 10
//  refer https://pkg.go.dev/fmt@go1.18.2
//   %v	the value in a default format
// %b	base 2
// %T Type of value (eg: for integer it shows int)
