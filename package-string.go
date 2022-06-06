package main

import (
	"fmt"
	"strings"
)

func main() {
	lower := strings.ToLower("HELLO WORLD")
	fmt.Println(lower)
	upper := strings.ToUpper("hello world")
	fmt.Println(upper)
	stringArray := []string{"I love", "go programming"}
	fmt.Println(stringArray)
	JoinStringArray := strings.Join(stringArray, " ")
	fmt.Println(JoinStringArray)
}

// convert the string to lowercase
// convert the string to uppercase
// create a string array
// join elements of array with space in between
