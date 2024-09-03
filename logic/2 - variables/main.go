package main

import "fmt"

var (
	nome string
	n1   int   = 1
	n2   int32 = 2
)

func main() {

	message := "Start Message"
	fmt.Println(message)

	var a, b, c = true, 2.3, "String"

	fmt.Println(a, b, c)

	var total int
	total++

	fmt.Println("Total:", total)

	var x, y = 10, 5

	fmt.Println("X:", x, "Y:", y)
	x, y = y, x
	fmt.Println("X:", x, "Y:", y)

	nome = "Arthur Silva"

	fmt.Println("Hello", nome, "!")
}
