package main

import "fmt"

func main() {
	cities := []string{"NYC", "COLOMBO", "LONDON", "TOKYO"}
	fmt.Println(cities)
	Addcity := append(cities, "AUKLAND")
	fmt.Println(Addcity)
	fmt.Println(cities)
	cities = append(cities, "LA")
	fmt.Println(cities)
}

// You can use the same method which we used to declare an Array to declare a Slice in Golang. But the difference is,
//  in slices you don’t have to mention the size for that particular slice. But in Arrays, we had to define its size when declaring.
//  Since a slice doesn’t have a specific size, we can append values to a slice in addition to changing its existing values.
//  You can append values to a slice using two different methods. You can append values to the original slice or creating another slice and append values without overwriting the original slice.
// you can use the built-in append() function to add values
