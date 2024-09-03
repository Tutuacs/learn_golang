package main

import "fmt"

type PessoaJuridica struct {
	RazaoSocial string
	cnpj        string
}

type Pessoa struct {
	Nome   string
	Idade  uint8
	Status bool
}

func (p Pessoa) String() string {
	return fmt.Sprintf("Olá, meu nome é %v e eu tenho %d anos. Printado do método de Pessoa", p.Nome, p.Idade)
}

type PessoaFisica struct {
	Pessoa
	Sobrenome string
	cpf       string
}

// Comentando o método de pessoa física ele procura o método de Pessoa
func (p PessoaFisica) String() string {
	return fmt.Sprintf("Olá, meu nome é %v e eu tenho %d anos. Printado do método de PessoaFísica", p.Nome, p.Idade)
}

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

	pessoa := pessoaFisica.Pessoa

	fmt.Println(pessoaFisica) // ele procura a função String() declarada acima
	fmt.Println(pessoa)

}
