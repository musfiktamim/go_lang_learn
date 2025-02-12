package main

import "fmt"

type OrderLocation string

const (
	recieved OrderLocation = "recieved"
	paid                   = "paid"
)

func orderPos(orderp OrderLocation) {
	fmt.Println(orderp)
}

func main() {
	orderPos(recieved)
}
