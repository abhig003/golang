package main
import "fmt"

func sum(a int, b int) int{
	return a+b
}

// what will happend if we will not give any type return then it will throw an err

// func sum1(a int, b int){
// 	return a+b
// }



// if the argumenet type is same then we dont need to pass argumenet type for evry argument we can ad in last the type of argument

func sum2(a,b int)int{

	return a+b
}

func multi_return() (string ,int){
 return "ajay",18

}


func array_return()[]string{
	return []string{"ajay","gupta","iisc"}
}

func main(){


	fmt.Println("here we will understand how we can define the function and their syntax, return mutiple value,  taking multiple type of argument")

	//

	//func function_name(argument argument_type) return type{

	//	return 
	//}


	fmt.Println(sum(3,5))
	//fmt.Println(sum1(3,5))// here its throwing error no value as used

	fmt.Println(sum2(3,6))
	val,_:=multi_return();// so here it will return the first return and the 2 return will ign ore by adding this "_"
	fmt.Println(val)

	_,val2:=multi_return()// so here it will return the 2nd argument
	fmt.Println(val2)

	list:=array_return();

	for i,val:=range list{
		fmt.Println("item of an array at index",i, "and their value",val)
	}

}