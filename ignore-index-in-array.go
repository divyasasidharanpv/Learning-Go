package main

import "fmt"

func main() {
	e := []int{1, 2, 3}

	for _, j := range e {
		fmt.Println(j)
	}
}

//ignoring index using "_"
//prints all elements in array using range function
//short variable declaration is required if variable is not declared
// _ is a variable and we cannot initialize twice. (we cannot use > for _, _ := range e)
