package main

import "fmt"

func main() {

	// Slice is an array with flexible size
	sayHi := []string{"Hello", "World", "🤪"}
	fmt.Println(sayHi)

	// Declare an array
	names := [5]string{
		"Nguyen",
		"Viet",
		"Duong",
		"Harry",
		"Na",
	}

	// Get element from the array ( x , y are slices )
	x := names[0:2] // [Nguyen Viet]
	y := names[2:5] // [Duong Harry Na]

	fmt.Println(x, y)

	y[0] = "🥳"
	fmt.Println(y) //  [🥳 Harry Na],
	// Changing the elements of a slice modifies the corresponding elements of its array.

	fmt.Println(names) // [Nguyen Viet 🥳 Harry Na]

	// length and capacity of array
	fmt.Println(len(names), cap(names))

	// Use append to add something new to the Slice.
	// Array could not because an array is fixed length.
	s := []string{}
	s = append(s, "Duong")
	fmt.Println(s)

}
