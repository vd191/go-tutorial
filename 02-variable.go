package main 

import "fmt"
import "math/cmplx"

// The var statement declares a list of variables
var batman, sup, flash bool

// A var declaration can include initializers, one per variable.
var power, strength int = 100, 400

// Basic types
var (
	ToBe bool = false
	MaxInt uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	var i int
	fmt.Println(batman, sup, flash, i)

	// The variable will take the type of the initializer.
	var fly, invisible = true, "no"

	// := statement can be used intead of var keyword. 
	joker := false
	fmt.Println(power, strength, fly, invisible, joker )

	// Test basic type
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)


	// Type conversions
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	// Constants cannot be declared using the := syntax.
	const sun = "☀️";

}




















