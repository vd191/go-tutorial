package main

import "fmt"

type Car struct {
	body, make, model string
}

/* A method is a function with a special receiver argument. */
func printDetail(c Car) string {
	message := fmt.Sprintf("Car details are: %s , %s , %s", c.make, c.model, c.body)
	return message
}

func main() {
	harryCar := Car{"SUV", "Volkswagen", "Tiguan"}
	fmt.Println(printDetail(harryCar))
}
