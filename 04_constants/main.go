package main

const age int = 30

func main() {
	//constant
	const name string = "musfikur rahman tamim"

	// grouping on constant variable
	const (
		host = "localhost"
		port = 89
	)
	//grouping on muteable variable
	var (
		names = "hello"
		ages  = 19
	)
	println(names)
	println(ages)
	println(name)
}
