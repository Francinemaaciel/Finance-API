# Finance API

API REST de controle financeiro pessoal desenvolvida em Go.

## 🚀 Tecnologias

- **Go** — linguagem principal
- **Chi** — roteador HTTP
- **PostgreSQL** — banco de dados
- **sqlx** — acesso ao banco
- **golang-migrate** — migrations
- **JWT** — autenticação
- **bcrypt** — criptografia de senhas
- **AwesomeAPI** — cotação do dólar em tempo real

## ⚙️ Como rodar

1. Clone o repositório
2. Crie o arquivo `.env`:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=sua_senha
DB_NAME=finance
JWT_SECRET=seu_secret
```
3. Rode as migrations:
```bash
migrate -path ./migrations -database "postgres://postgres:SUA_SENHA@localhost:5432/finance?sslmode=disable" up
```
4. Inicie o servidor:
```bash
go run cmd/main.go
```

### Com Docker
```bash
docker compose up -d
```
