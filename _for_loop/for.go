package main
import "fmt"


func main(){


	// for loop with init, condition, post statement

	//classic for loop
	for i:= 0; i<5;i++ {
		fmt.Println("Classic for loop iteration:", i)
	}

	// for loop as while loop
	j:=0
	for j<5 {
		fmt.Println("While-like for loop iteration:", j)
		j++
	}

	// infinite for loop with break
	k:=0
	for {
		if k>=5 {
			break
		}
		fmt.Println("Infinite for loop iteration:", k)
		k++
	}

	//infinete loop without any condition
	for {
		fmt.Println("This will run forever unless interrupted")
		break // added break to prevent actual infinite loop during execution
	}

	for i:=range 5{
		fmt.Println("iteration over for loop in range:", i)
	}

	// iteration over array using for loop
	var number =[5] int{10,20,30,40,50}
	for i,c:=range number{
		fmt.Println("Array element at index", i, "is", c)
	}

	for _,c:=range number{
		fmt.Println("Array element value is", c)
	}

	for i,_:=range number{
		fmt.Println("Array element index is", i)
	}

}