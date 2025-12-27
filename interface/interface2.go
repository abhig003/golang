// In this example, we are trying to understand why we need an interface
// and how it helps in solving complex real-world problems.
//
// Imagine we integrate a payment gateway into an existing project.
// Initially we may only support Razorpay, Paypal, or Slice.
// But later, the business may ask us to add more payment gateways.
// 
// Without interfaces, this becomes a problem because:
// 1. We have to modify main service code again and again.
// 2. Our code becomes tightly coupled with specific payment providers.
// 3. Changes become risky and difficult to maintain.
//
// To solve this, we use an INTERFACE.
// An interface acts like a CONTRACT.
// It defines what functions must exist (for example: Pay(amount int)).
// 
// Any payment gateway (Razorpay / Paypal / Slice / GooglePay / Stripe etc.)
// only needs to implement this interface. After that,
// it can be plugged into the system without changing the main business logic.
//
// So if tomorrow we want to add GooglePay,
// we just create a new struct and implement the Pay() method.
// We do not touch or modify existing working code.
// This makes the system:
// - scalable
// - maintainable
// - loosely coupled
// - easy to extend


package main

import "fmt"
import "flag"
import "strings"
type PaymentGateway interface{
	Pay(amount int)
}
// so if any payment method want to used the payment gateway they have need to only defined the Pay behavior in their method 
//

type MakePaypal struct{}
type MakeRazorpay struct{}
type MakeSlice struct{}
type MakeGoogle struct{}
// so here we have 4 type of payment method ,for using the interface contractor
//  we have need to defined the each payment Pay method function
func (paypal MakePaypal) Pay(amount int){
	fmt.Println("payment processed through paypal gateway",amount)
}
func (razorpay MakeRazorpay) Pay(amount int){
	fmt.Println("payment processed through razorpay",amount)
}

func (slice MakeSlice) Pay(amount int){
	fmt.Println("payment processed through slice",amount)
}
func (GooglePay MakeGoogle) Pay(amount int){
	fmt.Println("payment processed through google pay",amount)
}

func main(){

	fmt.Println("here we will undertand how we can used the interface to solve the complex problem ")

	// here we will take argument paymount mode type and amount

    paymentType := flag.String("payment", "slice", "payment gateway: razorpay | paypal | slice")
	amount := flag.Int("amount", 0, "payment amount")

	flag.Parse()


	var gateway PaymentGateway

	paypal:=MakePaypal{}
	slice:=MakeSlice{}
	googlepay:=MakeGoogle{}
	razorpay:=MakeRazorpay{}


	switch strings.ToLower(*paymentType){
	 case "paypal":
		gateway=paypal
     case "slice":
	    gateway	=slice
    case "razorpay":
		gateway=razorpay
    case "googlepay":
		gateway=googlepay				
	default:
	  fmt.Println("Invalid payment type! Use: paypal | razorpay | slice | google")
		return
	}

	gateway.Pay(*amount)
}