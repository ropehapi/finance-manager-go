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
	//TODO: Adicionar transactions nas operações
	_ = godotenv.Load()

	database := db.NewDatabase()

	accountRepo := repository.NewAccountRepository(database)
	transferRepo := repository.NewTransferRepository(database)
	paymentMethodRepo := repository.NewPaymentMethodRepository(database)

	accountService := service.NewAccountService(accountRepo)
	transferService := service.NewTransferService(transferRepo, accountRepo, paymentMethodRepo)
	//paymentMethodService := service.NewPaymentMethodService(paymentMethodRepo)

	accountHandler := handler.NewAccountHandler(accountService)
	transferHandler := handler.NewTransferHandler(transferService)
	//paymentMethodHandler := handler.NewPaymentMethodHandler(paymentMethodService)

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
	//paymentMethodHandler.RegisterRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
