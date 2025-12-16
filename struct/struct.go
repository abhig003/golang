package main

import (
	"fmt"
	"time"
)

// order struct

type customer struct {
	name  string
	phone string
}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func newOrder(id string, amount float32, status string) *order {
	
	// initial setup goes here...
 	myOrder := order{
 		id:     id,
 		amount: amount,
 		status: status,
 	}


	return &myOrder
 }

// // receiver type
 func (o *order) changeStatus(status string) {
 	o.status = status
 }

func (o order) getAmount() float32 {
	return o.amount
 }

func main() {
	// newCustomer := customer{
	// 	name:  "john",
	// 	phone: "1234567890",
	// }
	newOrder := order{
		id:     "1",
		amount: 30,
		status: "received",
		customer: customer{
			name:  "john",
			phone: "1234567890",
		},
	}

	newOrder.customer.name = "robin"
	

	// myOrder := newOrder("1", 30.50, "received")
	 //fmt.Println(myOrder.amount)
	// if you don't set any field, default value is zero value
	// int => 0, float => 0, string "", bool => false

	fmt.Println(newOrder)

		language := struct {
			name   string
			isGood bool
		}{"golang", true}

		fmt.Println(language)

	 myOrder3 := order{
		id:        "2",
	 	amount:    100,
 	status:    "delivered",
		createdAt: time.Now(),
	 }



	 myOrder3.changeStatus("confirmed")
	 fmt.Println(myOrder3)
	  myOrder3.createdAt = time.Now()
	  fmt.Println(myOrder3.status)

	 myOrder2 := order{
		id:        "2",
	 	amount:    100,
 	status:    "delivered",
		createdAt: time.Now(),
	 }

	 myOrder3.status = "paid"

	fmt.Println("Order struct", myOrder2)
	 fmt.Println("Order struct", myOrder3)
}