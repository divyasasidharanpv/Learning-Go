package main

import "fmt"

func main() {
	fmt.Println("Hello Go")

	funcA()
}
func funcA() {
	fmt.Println("Calling function a")
	controlflow()
}

func controlflow() {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

// // control flow
// 1. sequence
// 2. loop; iterative
// 3. conditional
