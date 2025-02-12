package main

import (
	"fmt"
)

func names() {
	fmt.Println("my name is musfik")
}

func add(a, b int) int {
	return a + b
}

func getLanguages() (string, string, string, int, bool) {
	return "golang", "cpp", "python", 3, false
}

func procces(fn func(a int) int) {
	fn(1)
}

func main() {
	names()
	var result int = add(100, 300)
	fmt.Println(result)
	lang1, lang2, lang3, _, _ := getLanguages()
	fmt.Println(getLanguages())
	fmt.Println(lang1, lang2, lang3)

	fn := func(a int) int { return a }

	procces(fn)
}
