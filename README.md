
# 💰 Finance Manager Go

API REST para gerenciamento de contas, escrita em Go utilizando GORM, PostgreSQL, Gin e Docker.

---

## 📦 Tecnologias utilizadas

- [Go 1.24.2](https://go.dev/)
- [PostgreSQL 16](https://www.postgresql.org/)
- [GORM](https://gorm.io/)
- [Gin](https://gin-gonic.com/)
- Docker + Docker Compose
- `jq` (para formatação de JSON no terminal)

---

## ⚙️ Pré-requisitos

Antes de começar, você precisa ter instalado:

- [Go 1.24.2](https://go.dev/dl/)
- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)

---

## 🚀 Rodando a aplicação

### 1. Clone o repositório

```bash
git clone https://github.com/ropehapi/finance-manager-go.git
cd finance-manager-go
```

### 2. Suba os containers com Docker Compose

```bash
make up-detach
```

> Ou, se não estiver usando Make:
>
> ```bash
> docker compose up -d
> ```

A aplicação estará disponível em:  
📍 `http://localhost:8080`

---

## 📂 Estrutura do projeto (Standard Layout)

```
finance-manager-go/
├── cmd/
│   └── api/                # Entrada da aplicação (main.go)
├── internal/
│   └── account/            # Domínio de contas (model, service, handler, etc)
├── pkg/
│   └── db/                 # Conexão com banco de dados
├── scripts/
│   └── wait-for-it.sh      # Script para aguardar o banco iniciar
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── README.md
```

---

## 🛠️ Principais comandos (Makefile)

```bash
make build         # Builda a imagem da API
make up            # Sobe os containers
make up-detach     # Sobe em background
make down          # Para os containers
make reset         # Para e remove volumes
make rebuild       # Builda sem cache
make health        # Testa se a API está online
make list-accounts # Lista contas existentes
make create-account # Cria uma conta de exemplo
```

---

## 📬 Usando a aplicação com um HTTP Client

### 📁 Importar coleção

1. Abra o Postman
2. Clique em **Import**
3. Selecione o arquivo `finance-manager-go.postman_collection.json`
4. Teste todos os endpoints de forma interativa

---

## 🐳 Observações sobre o Docker

- A API aguarda o PostgreSQL estar disponível antes de iniciar (`wait-for-it.sh`)
- O banco é inicializado com as credenciais:
  - **Usuário**: `postgres`
  - **Senha**: `postgres`
  - **Banco**: `accountdb`
- Os dados do banco são persistidos no volume `pgdata`
