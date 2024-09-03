package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
}

type Category struct {
	Nome string
	Pai  *Category // Precisa ser passada como ponteiro
}

func main() {
	names := []string{"Arthur", "Joao", "Pedro", "Matheus"}
	fmt.Println(names)

	nomes := append(names, "Mariana")

	fmt.Println(nomes)

	newNomes := make([]int, 3, 10) // make(tipo, "valores inicializados", "quantidade total esperada")

	fmt.Println(newNomes)

	mapAges := make(map[string]uint8)
	mapAges["Arthur"] = 21
	mapAges["Pedro"] = 22
	mapAges["Matheus"] = 23

	fmt.Println(mapAges)
	fmt.Println(mapAges["Arthur"])
	fmt.Println(mapAges["Pessoa não existente"]) // return the default key value "int = 0"

	pessoa := Pessoa{
		Nome:      "Arthur",
		Sobrenome: "Silva",
		Idade:     21,
		Status:    true,
	}

	fmt.Println(pessoa)
	fmt.Println(pessoa.Nome, pessoa.Idade)

	// Outra maneira não muito usada:
	pessoa2 := Pessoa{"Arthur", "Silva", 21, true}
	fmt.Println(pessoa2)

}
