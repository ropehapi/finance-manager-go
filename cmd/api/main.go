package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"

	"github.com/ropehapi/finance-manager-go/internal/handler"
	"github.com/ropehapi/finance-manager-go/internal/repository"
	"github.com/ropehapi/finance-manager-go/internal/service"
	"github.com/ropehapi/finance-manager-go/pkg/db"
)

func main() {
	_ = godotenv.Load() // carrega variáveis de ambiente do .env

	database := db.NewDatabase()

	// Repositórios
	accountRepo := repository.NewAccountRepository(database)
	transferRepo := repository.NewTransferRepository(database)

	// Services
	accountService := service.NewAccountService(accountRepo)
	transferService := service.NewTransferService(transferRepo, accountRepo)

	// Handlers
	accountHandler := handler.NewAccountHandler(accountService)
	transferHandler := handler.NewTransferHandler(transferService)

	// Router
	r := chi.NewRouter()
	accountHandler.RegisterRoutes(r)
	transferHandler.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
