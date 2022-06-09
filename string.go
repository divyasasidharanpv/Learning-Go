package main

import "fmt"

func main() {
	string := "Hello World"
	str := ` She said "Have 
	a
	good
	day"`
	fmt.Println(string, str)
}

// Creating a multiline string in Go is actually incredibly easy. Simply use the backtick ( ` ) character when declaring or assigning your string value.
//  Single line string using ( "")
// using back tick includes the double quotes . gives everything that's between those back ticks.
