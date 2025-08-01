# Variáveis
APP_NAME=finance-manager-go
API_PORT=8080

# Builda a imagem da API
build:
	docker compose build

# Sobe os containers em modo interativo
up:
	docker compose up

# Sobe os containers em modo detached
up-detach:
	docker compose up -d

# Para os containers
down:
	docker compose down

# Derruba tudo e remove volume do Postgres
reset:
	docker compose down -v

# Limpa cache e reconstroi
rebuild:
	docker compose build --no-cache

# Testa se a API está respondendo
health:
	curl -i http://localhost:$(API_PORT)/health

# Lista as contas cadastradas
list-accounts:
	curl -s http://localhost:$(API_PORT)/accounts | jq

# Cria uma nova conta
create-account:
	curl -X POST http://localhost:$(API_PORT)/accounts \
		-H "Content-Type: application/json" \
		-d '{"kind":"personal","currencyCode":"BRL","name":"Conta Corrente","balance":100000}' | jq

# Faz um commit
gcm:
	@if [ -z "$(m)" ]; then \
		echo "Erro: é necessário fornecer uma mensagem de commit."; \
		exit 1; \
	fi
	@git add . && git commit -m "$(m)" && git push