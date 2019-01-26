package main 

import "fmt"

// A struct is a collection of fields.
type Regtangle struct {
	// Struct fields 
	x,y int
}

var (
	r3 = Regtangle{1, 2}  // has type Regtangle
	r4 = Regtangle{x: 1}  // Y:0 is implicit
	r5 = Regtangle{}      // X:0 and Y:0
	p  = &Regtangle{1, 2} // has type *Regtangle
)

func main () {
	r1 := Regtangle{ 5, 10 }
	fmt.Println(r1)

	// Struct fields are accessed using a dot.
	r1.x = 8
	fmt.Println(r1.x)

	r2 := Regtangle{ x: 20, y: 50 }
	
	fmt.Println(area(r2))

	// Struct Literal 
	fmt.Println(r3, p, r4, r5)

}


func area (r Regtangle) int {
	return r.x * r.y
}