package main

import (
	"fmt"
	"time"
)

func numbers(n chan<- int) {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Escrito no channel: %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	close(n)
}

// func letters(done chan<- bool) {
// 	for l := 'a'; l < 'j'; l++ {
// 		fmt.Printf("%c ", l)
// 	}
// 	done <- true
// }

func main() {

	// ! Default channel
	// ! cn := make(chan int)

	// * Channel with buffer defined on 3 spaces
	cn := make(chan int, 3)
	go numbers(cn)

	for value := range cn {
		fmt.Printf("Lido do channel: %d\n", value)
		time.Sleep(time.Millisecond * 180)
	}

	// cl := make(chan bool)
	// go letters(cl)
	// <-cl

	fmt.Println("End of execution")
}
