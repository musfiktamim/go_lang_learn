package main

import (
	"fmt"
	"sync"
)

func task(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task doing on ", id)
}

func main() {
	var wg sync.WaitGroup
	for i := range 11 {
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()
}
