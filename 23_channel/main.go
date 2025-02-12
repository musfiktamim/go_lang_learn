package main

import (
	"fmt"
	"time"
)

// send
// func processNum(nums chan int) {
// 	for num := range nums {
// 		fmt.Println("processing number ", num)
// 		time.Sleep(time.Millisecond * 100)
// 	}

// }

// // reciev
// func recieved(nums chan int) {
// 	nums <- rand.Intn(100)
// }

// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	// defer func() { done <- false }()
// 	fmt.Println("proccessing...")
// 	time.Sleep(time.Second)
// }

func emailSender(emeilChan chan string, done chan bool) {
	defer func() { done <- true }()
	for emeilCha := range emeilChan {
		fmt.Println("email sended ", emeilCha)
		time.Sleep(time.Second)

	}
}

func main() {

	emailChan := make(chan string, 100)
	doneChan := make(chan bool)

	go emailSender(emailChan, doneChan)

	for i := 0; i < 5; i++ {
		fmt.Println(i)
		emailChan <- fmt.Sprintf("%d@example.com", i)
	}
	fmt.Println("completed")
	close(emailChan)
	<-doneChan
	// emailChan <- "hello@gmail.com"
	// emailChan <- "hello1@gmail.com"

	// fmt.Println(<-emailChan)
	// fmt.Println(<-emailChan)

	// done := make(chan bool)
	// for {
	// 	go task(done)
	// 	msg := <-done
	// 	fmt.Println(msg)
	// }

	// numchan := make(chan int)

	// go processNum(numchan)

	// for {
	// 	go recieved(numchan)
	// 	time.Sleep(time.Second)
	// }

	// mynum := make(chan string)

	// mynum <- "my name is musfik tamim"

	// msg := <-mynum

	// fmt.Println(msg)

}
