package main

import (
	"fmt"
	"time"
)

func numbers(done chan<- bool) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(time.Millisecond * 150)
	}
	done <- true
}

func letters(done chan<- bool) {
	for l := 'a'; l < 'j'; l++ {
		fmt.Printf("%c ", l)
		time.Sleep(time.Millisecond * 230)
	}
	done <- true
}

func main() {

	cn := make(chan bool)
	cl := make(chan bool)

	go numbers(cn)
	go letters(cl)

	<-cn
	<-cl

	fmt.Println("End of execution")
}
