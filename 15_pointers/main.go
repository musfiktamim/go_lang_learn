package main

import "fmt"

func changenumber(n *int) {
	*n = 5
}

func main() {
	var n = 1
	changenumber(&n)

	fmt.Println(n)
}
