# Go Product API (REST)

Uma API RESTful desenvolvida em **Go** para gerenciamento de produtos, utilizando uma arquitetura em camadas (Controller, UseCase, Repository) para garantir organização e facilidade de manutenção.

O projeto é totalmente containerizado com **Docker**, facilitando a configuração do ambiente e do banco de dados **PostgreSQL**.

-----------------------------------

## Créditos
Este projeto foi desenvolvido seguindo o tutorial do canal **Go Lab Tutoriais**: 
[Assista ao vídeo aqui](https://www.youtube.com/watch?v=3p4mpId_ZU8)

-----------------------------------

## Tecnologias Utilizadas

* **Linguagem:** [Go (Golang)](https://go.dev/) 1.23+
* **Framework Web:** [Gin Gonic](https://gin-gonic.com/)
* **Banco de Dados:** [PostgreSQL](https://www.postgresql.org/)
* **Driver DB:** [pq](https://github.com/lib/pq) (Pure Go Postgres driver)
* **Containerização:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)

-----------------------------------

## Arquitetura do Projeto

O projeto segue os princípios da **Clean Architecture**, dividido em:

1.  **Cmd:** Ponto de entrada da aplicação (`main.go`).
2.  **Controller:** Camada de interface que lida com as requisições HTTP e extração de parâmetros.
3.  **UseCase:** Camada de regras de negócio da aplicação.
4.  **Repository:** Camada de persistência que se comunica diretamente com o banco de dados.
5.  **Model:** Definição das estruturas de dados (Structs).
6.  **DB:** Configuração da conexão com o banco de dados.

-----------------------------------

## Como Executar o Projeto

Você precisará ter o **Docker** e o **Docker Compose** instalados em sua máquina.

1.  Clone o repositório:
    ```bash
    git clone [https://github.com/SEU_USUARIO/nome-do-seu-repositorio.git](https://github.com/SEU_USUARIO/nome-do-seu-repositorio.git)
    cd nome-do-seu-repositorio
    ```

2.  Suba os containers:
    ```bash
    docker compose up -d --build
    ```

A API estará disponível em: `http://localhost:8000`

-----------------------------------

## Documentação (Swagger)
A API possui documentação interativa via Swagger. Com a aplicação rodando, acesse:
[http://localhost:8000/swagger/index.html](http://localhost:8000/swagger/index.html)

-----------------------------------

## Endpoints da API

### Produtos

| Método | Endpoint | Descrição |
| :--- | :--- | :--- |
| **GET** | `/products` | Retorna a lista de todos os produtos. |
| **POST** | `/product` | Cria um novo produto (Envia JSON no Body). |
| **PUT** | `/product/:productId` | Atualiza um produto existente pelo ID. |
| **DELETE** | `/product/:productId` | Remove um produto do banco de dados. |

### Exemplo de JSON para Criação/Edição:
```json
{
  "product_name": "Batata",
  "price": 20
}
