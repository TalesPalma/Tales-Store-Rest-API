# Tales-Store-Rest-API

# Tales-Store-Rest-API

Bem-vindo ao **Tales-Store-Rest-API**! Este projeto é uma API RESTful desenvolvida em Go utilizando o framework Gin. A API é projetada para gerenciar e fornecer informações sobre produtos em uma loja fictícia.

## Índice

- [Requisitos](#requisitos)
- [Instalação](#instalação)
- [Uso](#uso)
- [Endpoints](#endpoints)
- [Testes](#testes)
- [Contribuição](#contribuição)
- [Licença](#licença)

## Requisitos

Antes de começar, verifique se você tem os seguintes softwares instalados:

- [Go](https://golang.org/doc/install) (v1.18 ou superior)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

## Instalação

1. Clone o repositório:

    ```bash
    git clone https://github.com/seu-usuario/tales-store-rest-api.git
    cd tales-store-rest-api
    ```

2. Instale as dependências:

    ```bash
    go mod tidy
    ```

## Uso

1. Para rodar a API, use o comando:

    ```bash
    go run main.go
    ```

    Por padrão, a API será iniciada na porta `8080`.

2. Você pode modificar a porta e outras configurações no arquivo `config/config.go`.

## Endpoints

Aqui estão alguns dos endpoints disponíveis na API:

- `GET /products` - Lista todos os produtos.
- `GET /products/:id` - Obtém um produto específico pelo ID.
- `POST /products` - Adiciona um novo produto.
- `PUT /products/:id` - Atualiza um produto existente pelo ID.
- `DELETE /products/:id` - Remove um produto pelo ID.

## Testes

Para executar os testes da API, use o comando:

```bash
go test ./...

