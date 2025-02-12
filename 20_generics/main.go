package main

import (
	"fmt"
)

func slicedPrint(sliced []int) {
	for _, slice := range sliced {
		fmt.Println(slice)
	}
}

func genericsslicedPrint[T comparable](sliced []T) {
	for _, slice := range sliced {
		fmt.Println(slice)
	}
}

type Stack struct {
	elements []int
}

type GenericsStack[T string | int] struct {
	elements []T
}

func main() {
	stack := Stack{elements: []int{12, 12, 23, 34, 45}}
	GenericsStack := GenericsStack[int]{elements: []int{12, 12, 12}}
	genericsslicedPrint(GenericsStack.elements)

	genericsslicedPrint(stack.elements)

	genericsslicedPrint([]int{1, 23, 34, 45, 56, 67, 12})
	genericsslicedPrint([]string{"1", "23", "34", "45", "56", "67", "12"})
}
