package main

import "fmt"

func main(){

	fmt.Println("here we will try to undertand , how we can define the map and how can iteratate over map, delete key, and how can clear the map")

	// make(map[type_key] type_val) make([string]int) so here the string is key and int is value

	var map1 =make(map[int] int);// if key is not exist then it will return 0 
	map1[1]=1;
	fmt.Println(map1)
	map1[2]=2
	fmt.Println(map1)

	val,bool_value:=map1[3]// it will return first arg as value and the 2nd argument as boolean if the lement exist then return true else false
	fmt.Println(val,bool_value)

	fmt.Println(len(map1))
	
	for i,val:=range map1{
		fmt.Println("key=",i ,"value=",val)
	}

    delete(map1,2)// his will delete the key element from the map
	delete (map1,1)
fmt.Println("after delete the key")
for i,val:=range map1{
		fmt.Println("key=",i ,"value=",val)
	}
	clear(map1)// it will delete the whole map
	fmt.Println(map1)
   map2:=map[int]string {1:"one",2:"two"}// intilsed without make key word 
   fmt.Println(map2)
  


}