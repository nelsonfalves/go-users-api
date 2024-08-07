# API Golang

## Descrição

Este projeto é uma API desenvolvida em Golang utilizando o banco de dados não relacional MongoDB. No momento, a API permite operações CRUD, e foi feita utilizando padrões de mercado modernos para garantir qualidade, escalabilidade e manutenibilidade.
O projeto ainda está em desenvolvimento e mais funcionalidades serão adicionadas.

## Tecnologias Utilizadas

- Golang
- MongoDB
- Docker
- Makefile

## Pré-requisitos

- Docker
- Docker Compose

## Instalação e Execução

### Passos Gerais

1. No terminal, para criar o binário executável e as imagens Docker, digite:

   make build-app

2. Em seguida, para subir a aplicação:

   make run-app

3. Agora já é possível realizar todas as operações no banco de dados:

#### Get All

- GET para `http://localhost:3000`

#### Get by ID

- GET para `http://localhost:3000/users/:id`

#### Create

- POST para `http://localhost:3000/users/create` junto de um body JSON no seguinte formato:

```
{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "password": "securepassword123"
}
```

#### Update

- PUT para `http://localhost:3000/users/update/:id` (junto de um body JSON no formato acima)

#### Delete

- DELETE para `http://localhost:3000/users/delete/:id`
