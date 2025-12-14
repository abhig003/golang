package main
import "fmt"
  
import "time"


func main(){
//switch statement example

 fmt.Println("here we will underatand how the swicth case work in go lang")

 // swicth condition{
 
 // case value1:
 //   //code to be executed if condition == value1
 // case value2:
 //   //code to be executed if condition == value2
 // default:
 //   //code to be executed if condition doesn't match any case
 // }

  // example 1: switch with integer values


  age:=18
  switch  {
  case age<12:
	fmt.Println("you are a child")
  case age>=13 && age<19:
	fmt.Println("you are a teenager")
  case age>=20 && age<60:
	fmt.Println("you are an adult")
  default:
	fmt.Println("you are a senior citizen")
	
  }

  // example 2: switch with string values

  day:="Wednesday"
  switch day{
  case "Monday":
	fmt.Println("Today is Monday")
  case "Tuesday":
	fmt.Println("Today is Tuesday")
  case "Wednesday":
	fmt.Println("Today is Wednesday")
  case "Thursday":
	fmt.Println("Today is Thursday")
  case "Friday":
	fmt.Println("Today is Friday")
  case "Saturday":
	fmt.Println("Today is Saturday")
  case "Sunday":
	fmt.Println("Today is Sunday")
  default:
	fmt.Println("Invalid day")
  }	

  // example 3: switch with fallthrough

  number:=2
  switch number{
  case 1:
	fmt.Println("One")
  case 2:
	fmt.Println("Two")
	fallthrough
  case 3:
	fmt.Println("Three")
  default:
	fmt.Println("Other number")
  }


  // example 4: switch with time package
  
  currentTime:=time.Now()

  fmt.Println("Current time is:", currentTime)
  switch currentTime.Weekday(){
  case time.Saturday, time.Sunday:
	fmt.Println("It's the weekend!")
  default:
	fmt.Println("It's a weekday.")
  }


}