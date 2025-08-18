# Variáveis
APP_NAME=finance-manager-go
API_HOST=http://localhost
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
	curl -i $(API_HOST):$(API_PORT)/health

# Lista as contas cadastradas
list-accounts:
	curl -s $(API_HOST):$(API_PORT)/accounts | jq

# Cria uma nova conta
create-account:
	curl --location '$(API_HOST):$(API_PORT)/accounts' \
    --header 'Content-Type: application/json' \
    --data '{"currencyCode": "BRL","name": "Conta corrente C6","balance": 50000}'

create-payment-method:
	curl --location '$(API_HOST):$(API_PORT)/payment-methods' \
    --header 'Content-Type: application/json' \
    --data '{"name": "Pix conta C6","type": "pix","accountId": "f853ea7d-adba-4691-b80b-5f2f3add8525"}'

# Faz um commit
gcm:
	@if [ -z "$(m)" ]; then \
		echo "Erro: é necessário fornecer uma mensagem de commit."; \
		exit 1; \
	fi
	@git add . && git commit -m "$(m)" && git push