package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	updatedAt time.Time
}

func constructOrder(id string, amount float32, status string) *order {
	constrOrder := order{
		id:     id,
		amount: amount,
		status: status,
	}

	return &constrOrder
}

// (*) deferencing

func (o *order) changeStatus(status string) (bool, string) {
	o.status = status
	return true, "status changes"
}

func (o order) getAmmount() float32 {
	var amount float32 = o.amount
	return amount
}

func main() {
	// var order order = order{id: "hey", amount: 12, status: "complete", createdAt: time.Now(), updatedAt: time.Now()}
	// order := order{
	// 	id:     "123213msdnjashd",
	// 	amount: 12321,
	// 	status: "recieved",
	// }
	// order.createdAt = time.Now()
	// order.updatedAt = time.Now()

	// order.changeStatus("paid")
	// fmt.Println(order)
	// fmt.Println("amount is ", order.getAmmount())
	orders := constructOrder("hey", 12, "recieved")
	fmt.Println(orders)

}
