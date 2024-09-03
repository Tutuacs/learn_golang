package main

import "fmt"

func main() {

	names := []string{"Arthur", "Joao", "Pedro", "Matheus"}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

	// Default FOR

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

	// Normal iteration

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

	// Default iteration for lists, get the range and value

	for i, nome := range names {
		fmt.Println(i, nome)
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

	// Default WHILE

	var i int

	for i < len(names) {
		fmt.Println(names[i])
		i++
	}

	// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

	// Default infinite listener

	for {
		fmt.Println("listening")
	}

}
