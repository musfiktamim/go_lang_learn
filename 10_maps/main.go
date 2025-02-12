package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string)
	m["name"] = "musfik"
	m["area"] = "backend"
	fmt.Println(m["name"])
	fmt.Println(m["age"])
	fmt.Println(m["area"])

	mi := make(map[string]int)
	mi["age"] = 17
	mi["ages"] = 18
	fmt.Println(mi["age"])
	fmt.Println(mi["ages"])
	delete(m, "ages")
	clear(mi)
	fmt.Println(mi["ages"])
	fmt.Println(len(mi))

	mb := make(map[string]bool)
	fmt.Println(mb)

	definedMaps := map[string]any{"age": 17, "name": "musfik"}
	definedMaps["age"] = 8

	k, ok := definedMaps["ages"]
	fmt.Println(ok)
	fmt.Println(k)

	fmt.Println(definedMaps)

	m1 := map[string]int{"price": 12, "phones": 5}
	m2 := map[string]int{"price": 12, "phones": 5}

	fmt.Println(maps.Equal(m1, m2))
}
