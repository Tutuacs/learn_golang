package main

import "fmt"

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

// Interface

type Document interface {
	Doc() string
}

// Interface Methods

func show(doc Document) {
	fmt.Println(doc)
	fmt.Println(doc.Doc())
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

// Structs

type PessoaJuridica struct {
	Pessoa
	RazaoSocial string
	cnpj        string
}

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

func (pf PessoaFisica) Doc() string {
	return fmt.Sprintf("meu cpf é: %s", pf.cpf)
}

func (pj PessoaJuridica) Doc() string {
	return fmt.Sprintf("meu cnpj é: %s", pj.cnpj)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - //

func main() {

	pessoaFisica := PessoaFisica{
		Pessoa: Pessoa{
			Nome:   "Arthur",
			Idade:  21,
			Status: true,
		},
		Sobrenome: "Silva",
		cpf:       "000.000.000.00",
	}

	pessoaJuridica := PessoaJuridica{
		Pessoa: Pessoa{
			Nome:   "Aprenda Golang",
			Idade:  0,
			Status: true,
		},
		RazaoSocial: "Temporin LTDA",
		cnpj:        "00.000.000/0000-00",
	}

	// pessoa := pessoaFisica.Pessoa

	fmt.Println(pessoaFisica)
	show(pessoaFisica)
	show(pessoaJuridica)
	// fmt.Println(pessoa)

}
