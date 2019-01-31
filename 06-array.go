package main

import "fmt"

func main() {
	// Declare an integer array with length of 10
	var a [4]int

	// Use square breaket to index an array
	a[3] = 7
	fmt.Println(a)

	// Declare an array
	names := [5]string{
		"Nguyen",
		"Viet",
		"Duong",
		"Harry",
		"Na",
	}

	fmt.Println(names) // [Nguyen Viet 🥳 Harry Na]

	// length and capacity of array
	fmt.Println(len(names), cap(names))

}
