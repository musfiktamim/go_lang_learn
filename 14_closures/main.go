package main

import "fmt"

func counter() func() int {
	var count int = 0

	return func() int {
		count += 1
		return count
	}
}

func main() {
	counterfun := counter()
	fmt.Println(counterfun())
	fmt.Println(counterfun())
	fmt.Println(counterfun())
}
