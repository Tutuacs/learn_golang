package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     uint8
	Status    bool
	cpf       string
}

type Category struct {
	Nome string
	Pai  *Category // Precisa ser passada como ponteiro
}

// Sem ponteiros é possivel apenas visualizar valores
func (category Category) IsChildren() bool {
	return category.Pai != nil
}

// É possivel alterar valores apenas com ponteiros
func (category *Category) SetPai(pai *Category) {
	category.Pai = pai
}

func main() {
	p := Pessoa{"Arthur", "Silva", 21, true, "000.000.000.00"}

	fmt.Println(p)

	cat := Category{Nome: "Category 1"}
	pai := Category{Nome: "Catgory PAI"}

	cat.SetPai(&pai)
	// cat.SetPai(&Category{Nome: "Pai"})

	if !cat.IsChildren() {
		fmt.Println("A categoria não tem um pai")
	} else {
		fmt.Println("A categoria tem um pai")
	}
}
