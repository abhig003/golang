// POINTERS IN GO
//
// A pointer is a variable that stores the memory address of another variable.
// In Go:
//   &variable  → gives the address (pointer)
//   *pointer   → gives the value stored at that address (dereferencing)
//
// Why pointers?
// - To access and modify the original value stored in memory
// - To avoid copying large data structures
// - Useful for performance and passing data by reference
//
// In this example:
//  We print the address of a variable
//  We access the value using dereferencing (*)
//  We change the value using a pointer

package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointer chapter")

	// Normal variable
	var name string = "ajay kumar gupta"

	// &name = memory address of 'name'
	fmt.Println("Address of variable 'name' is:", &name)

	// *(&name) means → go to that address and get the value
	fmt.Println("Value stored at that address is:", *(&name))

	// Store pointer in a variable
	ptr := &name

	// Change value using pointer
	*ptr = "abhi gupta"

	fmt.Println("After updating value using pointer:", name)
}
