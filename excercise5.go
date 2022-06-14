// 1. Create your own type of underlying type int.
// 2.Create a variable of your new type and identifier x using the var keyword (package level)
// Create a variable with identifier y and of type int.
// 3. In func main,
// a. Print out value of x
// b. print out type of variable x
// c. assign your value to x using = operator
// d. Print out value of variable x.
// Now use conversion to convert the type of the value stored in x to int and assign that value to y

package main

import (
	"fmt"
)

type aws int

var x aws
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 9
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
