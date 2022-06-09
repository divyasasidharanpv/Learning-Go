// //excercise3:
// 1. At package level scope, assign following values to variables
// a. for x assign 42
// b. for y assign "James Bond"
// c. for z assign true

// 2. In func main
// a. Use fmt.Sprintf to print all values to one string .
// Assign the return value of type string using short declaration operator to a variable  s
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z)
	fmt.Printf("%v\n", s)
}
