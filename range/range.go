package main

import "fmt"

func main(){


	//  we will iterate over an array, slice, map and number and strig , and understand how the range work

	var num [5]int// i will create an array of size 5 with 0 value [0,0,0,0,0]
	fmt.Println(num)
	var slice =make([]int,6,10)// it will create an slice of len 6 with capacity 10
	fmt.Println(slice)

	var Map=map[int]string{1:"ajay",2:"gupta"}
	fmt.Println(Map)


	var name= "ajay"
	fmt.Println(name)


	//range iteration
	//

	for i:=range 5{
		fmt.Println("iterating over the number range",i)
	}

	//itertaion over the array
	for i,val:=range num{// internally oits conevrt num=[a,b,c,d]
		fmt.Println("iterating over an array at index ",i , "and value on that index",val)
	}

	//iteration over slice
	for i,val:=range slice{
		fmt.Println("iterating over slice at index ",i , "and value on that index",val)
	}

	//itration over map
	for key,val:=range Map{
		fmt.Println("iteration over the map  and their key and value" ,key, val)
	}
}