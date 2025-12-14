package main

import "fmt"

import "slices"



func main(){
fmt.Println("we will try to understand how the slice work in golang")

// slice is a similar to the vector 
// struct slice{
//    *ptr,
//   len,
//   cap
  //}
// decalartion of slice 

num:=make([] int,2,10)// herewe have intilised the slices with int type with length 2 which contian 0 value at their index and the total capacity is 10
fmt.Println(num)//[0,0]

fmt.Println(len(num))//2
fmt.Println(cap(num))//10


// append the new element in slice

num=append(num,3,3,4,5);//it will return the updated slice
fmt.Println(num)

// empty slices is called nill 
val2:=make([]int,0);
fmt.Println(val2==nil)//this will return false becuse we have intilsed with zeo size
var s []int
fmt.Println(s == nil) // true


//copy slice into new slice

arr:=[]int{1,2,3,4,5}
fmt.Println(arr)


arr2 :=copy(make([]int, len(arr)),arr)//it will return the number of element copid in the new source

fmt.Println(arr2)
var arr3 =make([]int,len(arr))
copy(arr3,arr)//source and detination
fmt.Println(arr3)
//2d slices

var num4 =[][]int{{1,2},{3,4}}

fmt.Println(num4)
num5:=append(num4,[]int{1,2})
fmt.Println(num4,num5)


for i,val:=range num5{
	fmt.Println(num5[i],i,val)
}

fmt.Println(num5[:2])
fmt.Println(slices.Equal(num,num))// it will compare the two slice and and each element



}