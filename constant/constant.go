package main

import "fmt"
const age int = 25

func main() {

	fmt.Println("hello here we will try to undertand how the const works in  var and  is it a mutable or unmutable and how can we define  multiple const varible in  a single line")

	// const name type = value`
	const name string = "ajay"
	fmt.Println("the const variable is ", name)
	const name2 = "ajay"
	fmt.Println("the const variable is ", name2)

	// multiple const variable in a single line
	const (
		city    = "new york"
		country = "USA"
	)
	
	fmt.Println("the city is ", city)	
	fmt.Println("the country is ", country)

	fmt.Println("the age is ", age)
	//age=26 // this will give error because const variable is unmutable
	fmt.Println("the age after modification 1 is ", age)
	const age int = 26 // here we are trying to redefine the const variable which should  not be allowed but in the go their is conecpt o shadowing so this will create a new const variable with same name but different scope
	fmt.Println("the age after modification is ", age)
}
