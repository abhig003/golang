// here we will try to create an interface and wy  we need interface and how it will solve the complex problem
// so lets create an exapmle where we have need to integarte the payment gateway in th existng project and later we have need 
//to add other payment IN gateway so what will be challenge comes to integarte it in exiting gateway and how can we overcomes
//by using interface as a contractor

package main

import "fmt"


//razorpay, paypal, and slice payment


type MakeRazorpay struct{

}

type MakePaypal struct{

}

type MakeSlice struct {
	
}

// here we will define the payment function for evry payment

func (razorpay MakeRazorpay) payment(amount int){
   fmt.Println("payment throgh razorpay " ,amount)
}

func (payapl MakePaypal ) payment(amount int){
	fmt.Println("paymnet through paypal",amount)
}
func (slice MakeSlice) payment(amount int){
	fmt.Println("paymnet through slice",amount)
}
func main(){
	fmt.Println("so here we will integarte the payment gateway")

	slice:=MakeSlice();
	razorpay:=MakeRazorpay();
	paypal:=MakePaypal();

	slice.payment(30)
	razorpay.payment(40)
	paypal.payment(50)

}

