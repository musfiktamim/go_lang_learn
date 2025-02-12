package main

import (
	"fmt"
	"time"
)

type user struct {
	username string
	email    string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time
	user      user
}

func (o *order) setAmmount(amount float32) {
	o.amount = amount
}

func main() {
	var myorder order = order{id: "asdjksad", status: "recieved", createdAt: time.Now(), user: user{email: "musfiktamim@gmail.com"}}
	myorder.user.username = "musfiktamim"
	// myorder.user.email = "musfiktamim@gmail.com"
	myorder.setAmmount(20.0)
	fmt.Println(myorder)
}
