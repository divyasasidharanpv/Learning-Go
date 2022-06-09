package main

import "fmt"

var a int

type number int

var b number

func main() {
	a = 43
	b = 23
	a = int(b)

	fmt.Printf("Type of b is %T\n", b)
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("value of b is", b)
	fmt.Println("value of a is", a)
}

// a is of type int and b is of type number; inorder to assign a = b,  here value of b (type number) is converted to int and  is assigned to a.
// a is declared type int and it can store only int values
// value of b ( type number ) is  converted to int not type of b.
// value stored in b is of type number.
