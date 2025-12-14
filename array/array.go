package main

import "fmt"

func main() {

	fmt.Println("here we will understand how we can iterate over array, how we can define the array and what is deffrence between slice and array")

	// syntax=: [size of array] [type of elmenet]
	num := [5]int{} // and if we will not add {} bracket then it will throgh error becuse the ":=" this the shortcut way to define the varibale and assign the value at atime and if we will just instilsed then it will throgh error so instes of ":=" we can intilised with var"
	var num1 [5]int // thsi is the way we can intlised the array with 0 value
	// [0,0,0,0,0]
	fmt.Println(num)
	fmt.Println(num1)

	//len of array
	fmt.Println("the length of an array", cap(num), len(num))

	num2 := [5]int{1, 2, 3}
	fmt.Println(num2)

	num3 := [10]int{3: 400, 500}
	fmt.Println("here we can asign the value at from randowm index", num3)

	arr := []int{1, 2, 3, 4, 5}
	arr[0] = 2
	fmt.Println(arr[0:4]) // it will slice the array from starting to last_index-1, here the last_index is the last aragument , here we have given 4 as 2nd argument so this will return till 0,1,2,3
	// and if we will not give any first argument then it will conside as 0 by deafult
	fmt.Println(arr[:4])

	for i, v := range arr {
		fmt.Println("index and their value ", i, v)
	}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	
}
