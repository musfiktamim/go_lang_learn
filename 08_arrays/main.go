package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	nums[1] = 2
	nums[2] = 3
	nums[3] = 4

	var vals [4]bool

	var str [4]string

	names := [5]int{1, 2, 3, 4, 5}
	names2d := [2][2]int{{1, 2}, {3, 4}}

	fmt.Println(len(nums))
	fmt.Println(nums[0])
	fmt.Println(nums[1])
	fmt.Println(nums[2])
	fmt.Println(nums[3])
	fmt.Println(nums)
	fmt.Println(vals)
	fmt.Println(str)
	fmt.Println(names)
	fmt.Println(names2d)
	// use array if that prdicatable -- for memory optimization,constant time access

}
