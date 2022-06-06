package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(25))
	fmt.Println(math.Cbrt(64))
	fmt.Println(math.Mod(5, 2))
	fmt.Println(math.Max(11, 12))
	var value float64 = math.Min(22, 11)
	fmt.Printf("%f", value)
}

// find the square root
// find the cube root
// find the remainder
// find the maximum number
// find the minimum number
