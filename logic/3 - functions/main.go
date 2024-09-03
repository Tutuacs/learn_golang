package main

import (
	"fmt"
	"strconv"
)

func hello(nome string) {
	fmt.Println("Hello", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) int {
	i, _ := strconv.Atoi(b)

	return a + i
}

func main() {
	hello("Arthur Silva")

	fmt.Println("total:", sum(10, 3))
	fmt.Println("total:", convertAndSum(10, "8"))

}
