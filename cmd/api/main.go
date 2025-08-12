package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/ropehapi/finance-manager-go/internal/handler"
	"github.com/ropehapi/finance-manager-go/internal/repository"
	"github.com/ropehapi/finance-manager-go/internal/service"
	"github.com/ropehapi/finance-manager-go/pkg/db"
)

func main() {
	//TODO: Adicionar transactions nas operações: Aparentemente já é feito por padrão no gORM
	//TODO: Tornar filtragens dos endpoints GET genéricos
	//TODO: Remover category como struct e ser apenas string
	//TODO: Fazer relacionamentos no banco
	//TODO: Adicionar testes unitários
	_ = godotenv.Load()

	database := db.NewDatabase()

	accountRepo := repository.NewAccountRepository(database)
	transferRepo := repository.NewTransferRepository(database)
	paymentMethodRepo := repository.NewPaymentMethodRepository(database)

	accountService := service.NewAccountService(accountRepo)
	transferService := service.NewTransferService(transferRepo, accountRepo, paymentMethodRepo)
	paymentMethodService := service.NewPaymentMethodService(paymentMethodRepo)

	accountHandler := handler.NewAccountHandler(accountService)
	transferHandler := handler.NewTransferHandler(transferService)
	paymentMethodHandler := handler.NewPaymentMethodHandler(paymentMethodService)

	r := gin.Default()
	account := r.Group("/accounts")
	account.POST("/", accountHandler.Create)
	account.GET("/", accountHandler.GetAll)
	account.GET("/:id", accountHandler.GetByID)
	account.PUT("/:id", accountHandler.Update)
	account.DELETE("/:id", accountHandler.Delete)

	transfer := r.Group("/transfers")
	transfer.POST("/cashin", transferHandler.Cashin)
	transfer.POST("/cashout", transferHandler.Cashout)
	transfer.GET("/", transferHandler.GetAll)
	transfer.GET("/:id", transferHandler.GetByID)
	transfer.DELETE("/:id", transferHandler.Delete)

	paymentMethod := r.Group("/payment-methods")
	paymentMethod.POST("/", paymentMethodHandler.Create)
	paymentMethod.GET("/", paymentMethodHandler.GetAll)
	paymentMethod.GET("/:id", paymentMethodHandler.GetByID)
	paymentMethod.PUT("/:id", paymentMethodHandler.Update)
	paymentMethod.DELETE("/:id", paymentMethodHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
