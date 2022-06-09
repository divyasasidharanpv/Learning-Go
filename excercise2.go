// Use var to declare 3 variables. Do not assign values to the variables. Use the following identifiers
// a. identifier x of type int
// b. identifier y of type string
// c. identifier z of type bool
// In func mainn
// Print out values for each identifier
// the compiler assigned values to the variables are called ?

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("The compiler assigned values to the variables are called zero values")
}
