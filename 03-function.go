package main

import "fmt"
// Or can also write multiple import statement. 
import "math"

func main () {
	// Print pi number
	// a name is exported if it begins with a capital letter.
	fmt.Println(math.Pi)

	// Call add function 
	fmt.Println(add(1,1))

	// Retrieve 2 returns form function. 
	a,b := swap("😝", "🧐")
	//Call swap function 
	fmt.Println(a,b)

	// Named return values
	fmt.Println(split(20))

}

// func [function name] ( [argument] ) [return type] {}
func add(x int, y int) int { // (x , y int) : Functions continued
	return x + y
}

// Function return multiple results.
func swap(a, b string) ( string, string ) {
	return b, a
}

// Go's return values may be named. 
// If so, they are treated as variables defined at the top of the function.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	
	//A return statement without arguments returns the named return values. 
	// This is known as a "naked" return.
	return 
} 

