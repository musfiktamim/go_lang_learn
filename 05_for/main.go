package main

func main() {
	// for i := 1; i <= 100; i++ {
	// 	println("hello world", i)
	// }
	var i int = 1
	for i <= 100 {
		println(i)
		i++
	}

	for j := 0; j <= 100; j++ {
		println(j)
	}

	// for {
	// 	println("hello world", i)
	// 	i++
	// }
	for {
		println("hello worlds")
		break
	}
	for {
		println("hello worlds")
		break
		continue
	}

	for k := range 11 {
		println(k)
	}
}
