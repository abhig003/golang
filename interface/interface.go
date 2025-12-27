// here we will try to create an interface and wy  we need interface and how it will solve the complex problem
// so lets create an exapmle where we have need to integarte the payment gateway in th existng project and later we have need 
//to add other payment IN gateway so what will be challenge comes to integarte it in exiting gateway and how can we overcomes
//by using interface as a contractor

package main

import (
	"flag"
	"fmt"
	"strings"
)

type MakeRazorpay struct{}
type MakePaypal struct{}
type MakeSlice struct{}

func (razorpay MakeRazorpay) payment(amount int) {
	fmt.Println("payment through razorpay", amount)
}

func (paypal MakePaypal) payment(amount int) {
	fmt.Println("payment through paypal", amount)
}

func (slice MakeSlice) payment(amount int) {
	fmt.Println("payment through slice", amount)
}

func main() {
	fmt.Println("Payment Gateway Example with Flags")

	
	paymentType := flag.String("payment", "slice", "payment gateway: razorpay | paypal | slice")
	amount := flag.Int("amount", 0, "payment amount")

	
	flag.Parse()


	razorpay := MakeRazorpay{}
	paypal := MakePaypal{}
	slice := MakeSlice{}

	switch strings.ToLower(*paymentType) {
	case "razorpay":
		razorpay.payment(*amount)

	case "paypal":
		paypal.payment(*amount)

	case "slice":
		slice.payment(*amount)

	default:
		fmt.Println("Invalid payment type. Use: razorpay | paypal | slice")
	}
}


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
