package main 

import "fmt"
import "runtime"
import "time"

func main () {
	// defer statement
	defer fmt.Println("This will be executed until the surrounding function returns.")

	sum := 0;

	// Go has only one looping construct, the for loop.
	for i := 0; i <10; i++ {
		sum += i
	}

	fmt.Println(sum)

	// For continued. The init and post statements are optional. Can drop the semicolons
	multiply := 2;
	for multiply < 20 {
		multiply *= multiply
	}

	fmt.Println(multiply)

	// If statement
	fmt.Println( "sum", check(sum), "and multiply", check(multiply))

	// Switch case
	fmt.Println("Go runs on") 
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("MAC OS")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)
	}

	// Switch evaluation order
	// comparing the case expression. Stop if case expression equal switch expression
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}


	// Switch with no condition
	//This construct can be a clean way to write long if-then-else chains.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

}


func check(x int) string {

	// If statement
	if x < 100 {
		return " less than 100"
	}

	return "more than 100"
}
