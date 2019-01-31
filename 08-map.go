package main

import "fmt"

func main() {
	// map is a collection of keys and values
	cars := make(map[string]string)

	cars["Register number"] = "CN44F"
	cars["Year"] = "2012"
	cars["Make"] = "Volkswagen"
	cars["model"] = "Tiguan"

	fmt.Println(cars["Year"])

	delete(cars, "Tiguan")

	fmt.Println(cars)
}
