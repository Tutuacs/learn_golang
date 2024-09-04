<p align="center">
  <a href="" rel="noopener">
    <img width=200px height=200px src="https://miro.medium.com/v2/resize:fit:1000/0*YISbBYJg5hkJGcQd.png" alt="Project logo">
  </a>
</p>

<h3 align="center">Go Todo API</h3>

<div align="center">

[![Status](https://img.shields.io/badge/status-active-success.svg)]()
[![GitHub Issues](https://img.shields.io/github/issues/Tutuacs/learn_golang.svg)](https://github.com/Tutuacs/learn_golang/issues)
[![GitHub Pull Requests](https://img.shields.io/github/issues-pr/Tutuacs/learn_golang.svg)](https://github.com/Tutuacs/learn_golang/pulls)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

</div>

---

<p align="center"> Uma API simples de gerenciamento de tarefas (Todo) construída em Go (Golang).
    <br> 
</p>

## 📝 Table of Contents

- [Sobre](#about)
- [Requisitos](#requisitos)
- [Instalação](#instalacao)
- [Startando a Aplicação](#startando-a-aplicacao)
- [Uso](#usage)
- [Deployment](#deployment)
- [Built Using](#built_using)
- [Autores](#authors)
- [Reconhecimentos](#acknowledgement)

## 🧐 Sobre <a name = "about"></a>

Este projeto é uma API simples de gerenciamento de tarefas (Todo) construída em Go (Golang) usando `chi` para roteamento e PostgreSQL como banco de dados.

## 💻 Requisitos <a name = "requisitos"></a>

- Go 1.16 ou superior
- PostgreSQL
- Git

## 🛠 Instalação <a name = "instalacao"></a>

1. Clone o repositório:
   ```sh
   git clone https://github.com/Tutuacs/learn_golang.git
   cd learn_golang
   ```

2. Instale as dependências:
   ```sh
   go mod tidy
   ```

3. Configure o banco de dados PostgreSQL com as credenciais corretas. As configurações estão no arquivo `config.toml`:
   ```toml
   [api]
   port = "9000"

   [database]
   host = "localhost"
   port = "5432"
   user = "user_todo"
   pass = "1122"
   name = "api_todo"
   ```

## 🚀 Startando a Aplicação <a name = "startando-a-aplicacao"></a>

Para iniciar a aplicação, dentro da pasta `service`, execute o seguinte comando:
```sh
go run main.go
```

## 🎈 Uso <a name="usage"></a>

Acesse a API via *`http://localhost:9000`* após iniciar o servidor.

## 🚀 Deployment <a name = "deployment"></a>

Para fazer o deploy da aplicação em um ambiente de produção, configure o PostgreSQL em um servidor e inicie a aplicação com as mesmas configurações utilizadas em desenvolvimento.

## ⛏️ Built Using <a name = "built_using"></a>

- [Go](https://golang.org/) - Linguagem de Programação
- [Chi](https://github.com/go-chi/chi) - Framework de Roteamento
- [PostgreSQL](https://www.postgresql.org/) - Banco de Dados

## ✍️ Autores <a name = "authors"></a>

- [Arthur Silva](https://github.com/seu-usuario) - Desenvolvimento Inicial

## 🎉 Reconhecimentos <a name = "acknowledgement"></a>

- Agradecimento a todos que contribuíram para o desenvolvimento deste projeto.
