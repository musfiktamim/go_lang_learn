package main

import "fmt"

func sum(nums ...int) int {
	var Sum int = 0

	for _, num := range nums {
		Sum += num
	}

	return Sum
}

func main() {
	var slices = []int{12, 23, 34, 56, 67, 123, 34, 12}
	fmt.Println(sum(slices...))

	fmt.Println("hello world")
}
