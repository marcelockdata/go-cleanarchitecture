# Clean Architecture API

Esta aplicação é uma API RESTful construída em Go, seguindo os princípios de Clean Architecture. Ela oferece autenticação de usuários e gerenciamento de produtos, com documentação automática via Swagger.

## Principais Tecnologias

- **Go** (>= 1.23.4)
- **SQLite** (persistência local)
- **Docker & Docker Compose**
- **Swagger** (documentação automática)
- **Chi** (roteador HTTP)

## Como rodar com Docker

1. **Pré-requisitos**
   - Docker e Docker Compose instalados

2. **Build e execução**
   ```bash
   docker-compose up --build
   ```

3. **Acesse a API**
   - Swagger: [http://localhost:3000/swagger/index.html](http://localhost:3000/swagger/index.html)
   - Endpoints principais:
     - `POST /auth` — Autenticação de usuário
     - `GET /products` — Listagem de produtos
     - `POST /products` — Criação de produto
     - `GET /products/{id}` — Detalhe de produto
     - `PUT /products/{id}` — Atualização de produto
     - `DELETE /products/{id}` — Remoção de produto

## Estrutura do Projeto

```
├── cmd/
│   └── main.go
├── internal/
│   ├── handler/
│   ├── middleware/
│   ├── repository/
│   └── usecase/
├── docs/
│   └── swagger gerado pelo swag
├── products.db
├── Dockerfile
├── docker-compose.yml
└── README.md
```

## Observações

- O banco de dados `products.db` é persistido localmente e montado no container Docker.
- O Swagger é gerado automaticamente via [swag](https://github.com/swaggo/swag).  
  Para atualizar a documentação, rode:
  ```bash
  swag init --output ./docs --generalInfo ./cmd/main.go
  ```
- O driver SQLite exige CGO, por isso o Dockerfile usa uma imagem baseada em Debian e instala as dependências necessárias.

## Autenticação

- As rotas de produtos exigem autenticação via token JWT.
- Use o endpoint `/auth` para obter um token.

## Licença

Apache 2.0

---

**Dúvidas ou sugestões?**  
Abra uma issue ou envie um