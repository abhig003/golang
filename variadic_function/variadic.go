package main

import "fmt"

func sum(nums ...int){

		fmt.Println("user have pass the argument",nums)
	}

func main(){

	//we ill understand what is an variadic function and how is similar to the rest opearor and spread opeartor in js

	fmt.Println("welcome to the new conecpt of variadic function")

	// so the conecpt comes when we does know the actually the total number of aargument we have need to pass sometimes less ,sometime more,so does not know exactly the argument so for that we can used teh avriadic function so 
	//so in the variadic function we can pass any number of argument but when we define the function , we have need to follow some rules
	//here is the syntx

	// func variadic_function(argument_variable_name ...type_of argument){
	// 	// this will return the array of an argument and we can perform any opeartion
	// }


    // func sum(nums ...int){// internally its become []int

	// 	fmt.Println("user have pass the argument",nums)
	// }
	//sum(1,2,3,4) it will print [1,2,3,4]


	sum(1,2,3,4,)

	num:=[]int{1,2,3,4,5}
	fmt.Println(num)

	sum(num ...)
	///fmt.Println(...num) this will not work

	// varidiac function can also be used for appending 2 slices
a := []int{1, 2}
b := []int{3, 4}

a = append(a, b...)

    
	//
}
