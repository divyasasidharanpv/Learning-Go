package main

import "fmt"

func main() {
	e := []int{1, 2, 3, 4, 5}
	var i, j int
	for i, j = range e {
		fmt.Println(i, j)
	}
}

//print all elements in array using range function. For loop has a default range function.
// in this case, index is stored in i and elements in j.
