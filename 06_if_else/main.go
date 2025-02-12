package main

func main() {
	var age int = 41
	if age >= 18 && age <= 40 {
		println("you are adult")
	} else if age > 10 && age < 18 {
		println("you are teen")
	} else if age > 40 {
		println("you are expired")
	}

	if age > 10 || age < 10 {
		println("hey")
	}
}
