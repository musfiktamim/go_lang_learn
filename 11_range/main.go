package main

import "fmt"

func main() {
	nums := []int{5, 2, 3, 4, 5, 6}
	for i := 0; i < len(nums); i++ {
		// fmt.Println(nums[i])
	}
	sum := 0
	for v, num := range nums {
		sum += num
		fmt.Println(num, v)
	}

	fmt.Println(sum)

	mapss := map[string]string{"fname": "musfikur rahman", "lname": "tamim"}

	for key, value := range mapss {
		fmt.Println(key, value)
	}

	names := "hey whats app guys"

	for i, c := range names {
		fmt.Println(i, " : ", c, " = ", string(c))
	}

}
