package main 

import "fmt"

func main () {
	// Declare an integer array with length of 10 
	var a [4]int

	// Use square breaket to index an array 
	a[3] = 7
	fmt.Println(a)

	// Slice is an array with flexible size 
	b := []string{"Hello", "World", "🤪"}
	fmt.Println(b)

	// Slices are like references to arrays
	names := [5]string {
		"Nguyen", 
		"Viet", 
		"Duong", 
		"Harry", 
		"Na",
	} // fixed array

	x := names[0:2]		// [Nguyen Viet]
	y := names[2:5]  	// [Duong Harry Na]

	fmt.Println(x, y)

	y[0] = "🥳" 
	fmt.Println(y) //  [🥳 Harry Na], 
	// Changing the elements of a slice modifies the corresponding elements of its array.

	fmt.Println(names) // [Nguyen Viet 🥳 Harry Na]

	// length and capacity of array
	fmt.Println(len(names), cap(names))

	// Use append to add something new to the array. Slice only. 
	s := []string{}
	s = append(s, "Duong")
	fmt.Println(s)





}