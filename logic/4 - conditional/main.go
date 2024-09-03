package main

import (
	"fmt"
	"log"
	"os"
)

var (
	cara, coroa int
)

func jogarMoeda(lado string) {
	switch lado {
	case "cara":
		cara++
	case "coroa":
		coroa++
	default:
		fmt.Println("Caiu em pé")
	}
}

func main() {
	a, b := 10, 5

	// if a > b {
	// 	fmt.Println("A é maior que B")
	// } else if a < b {
	// 	fmt.Println("A é menor que B")
	// } else {
	// 	fmt.Println("A é igual a B")
	// }

	switch {
	case a > b:
		fmt.Println("A é maior que B")
	case a < b:
		fmt.Println("A é menor que B")
	default:
		fmt.Println("A é igual a B")

	}

	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}

	data := make([]byte, 100)

	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))

}
