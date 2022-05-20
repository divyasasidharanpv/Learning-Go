package main

import "fmt"

func main() {
	array := [4]int{8, 9, 10}
	fmt.Println("Size of array is", len(array), array)
	array[2] = 11
	fmt.Println("array value changed from 10 to", array[2], array)

}

//When you declare an Array, you need to clearly define its size and when you initialize the Array, you can assign values up to that size.
// // array size is 4. Here, you can add one more value and cannot exceed 4
// When you need to change Array values, you can use the relevant index and change its value accordingly.
// Once you define and set up Array values, you can easily print them using the “fmt”.
// fmt.Println(array)
