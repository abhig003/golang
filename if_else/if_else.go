package main
import "fmt"


func main(){
 // if-else statement example
 age := 20

 if age>18{
	fmt.Println("you are eligible for vote")
 } else{
	fmt.Println("you are not eligible for vote")
 }

  //inline variable declaration in if-else
 if score:=85; score>=50{
	fmt.Println("you have passed the exam with score:", score)
 } else{
	fmt.Println("you have failed the exam with score:", score)
 }
// if-else if-else ladder
 marks:=75
 if marks>=90{
	fmt.Println("Grade A")
 } else if marks>=75{
	fmt.Println("Grade B")
 } else if marks>=50{
	fmt.Println("Grade C")
 } else{
	fmt.Println("Grade F")
 }
}