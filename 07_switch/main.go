package main

func main() {
	//simple switch
	// var i int = 100
	// switch i {
	// case 1:
	// 	println("the number is one")
	// 	break // automaticaly break you does'nt have to define it
	// case 100:
	// 	println("the number is hundrad")
	// 	break
	// default:
	// 	println("other number is here")
	// }
	//mutliple switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	println("its weekend")
	// default:
	// 	println("its work day")
	// }

	//type switch
	// i := 10
	whois := func(i interface{}) {
		switch i.(type) {
		case int:
			println("integer")
		case string:
			println("hello world")
		}
	}

	whois(12387)

}
