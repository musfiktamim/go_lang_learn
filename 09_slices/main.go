package main

import (
	"fmt"
	"slices"
)

func main() {
	var num []int
	fmt.Println(num)
	fmt.Println(num == nil)
	fmt.Println(len(num))

	var nums = make([]int, 0, 2)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 7)
	nums = append(nums, 8)
	fmt.Println(cap(nums))
	fmt.Println(nums)

	var sliceds = []int{1, 2, 3, 4, 5, 6}
	fmt.Println(cap(sliceds))
	fmt.Println(len(sliceds))
	fmt.Println(sliceds)

	copyedsliced1 := make([]int, 0, 1)
	copyedsliced1 = append(copyedsliced1, 10)
	copyedsliced1 = append(copyedsliced1, 20)
	copyedsliced1 = append(copyedsliced1, 30)
	copyedsliced1 = append(copyedsliced1, 40)
	copyedsliced1 = append(copyedsliced1, 50)
	copyedsliced2 := make([]int, len(copyedsliced1), cap(copyedsliced1))

	copy(copyedsliced2, copyedsliced1)

	fmt.Println(copyedsliced1, copyedsliced2)

	// sliced oparetor
	var numsop = []int{1, 2, 3, 4, 5}
	fmt.Println(numsop[:])

	// slice
	var slice1 = []int{1, 2, 3}
	var slice2 = []int{1, 2, 4}
	fmt.Println(slices.Equal(slice1, slice2))

}
