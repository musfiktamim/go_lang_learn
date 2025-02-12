package main

import "fmt"

type paymanter interface {
	pay(amount float32)
}

type payments struct {
	payment paymanter
}

func (pts *payments) makePayment(amount float32) {
	pts.payment.pay(amount)
}

type bkash struct{}
type nagad struct{}

func (b bkash) pay(amount float32) {
	fmt.Println("payments by bkash")
}

func (n nagad) pay(amount float32) {
	fmt.Println("payments by nagad")
}

func main() {
	// var bkash bkash = bkash{}
	var nagad nagad = nagad{}
	var payments payments = payments{payment: nagad}
	payments.makePayment(123)
}
