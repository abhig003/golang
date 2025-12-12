package main

import "fmt"
func main(){

fmt.Println("hello world")
// in golang we can define the variable 3 types
// var name type = value
//var name=value(here the type will be infered from the value)
//name:=value(here the type will be infered from the value and also we can use this only inside the function)



var name string="ajay"
var name2="ajay"
name3:="ajay"

fmt.Println("the first way to define variable",name)
fmt.Println("the second way to define variable",name2)
fmt.Println("the third way to define variable",name3)


}