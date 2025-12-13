package main

import "fmt"

const age int = 25

func main() {
	
	const name string = "ajay"
	fmt.Println("the const variable is ", name)

	// defining multiple constants in a single block
	const (
		city    = "new york"
		country = "USA"
	)
	fmt.Println("the city is ", city)
	fmt.Println("the country is ", country)

	fmt.Println("the age is ", age)

	// age = 26 // not allowed: constants are immutable
	fmt.Println("the age after modification 1 is ", age)

	// this is not  redefinition or modification.
	// Go allows shadowing: a new const named age is created
	// inside this function scope, which hides the outer age.
	const age int = 26

	fmt.Println("the age after modification is ", age)
}
