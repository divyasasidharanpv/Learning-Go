package main

import "fmt"

func main() {
	//  define array
	a := []int{1, 2, 3, 4, 5, 6}
	// loop over array
	for x := 0; x < len(a); x++ {
		fmt.Println("Character:", a[x])
	}
}
