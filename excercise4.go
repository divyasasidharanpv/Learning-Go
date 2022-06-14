// 1. Create your own type of underlying type int.
// 2.Create a variable of your new type and identifier x using the var keyword
// 3. In func main,
// a. Print out value of x
// b. print out type of variable x
// c. assign 42 to variable x using the = operator
// d. Print out value of variable x.

package main

import "fmt"

type aws int

var x aws

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
