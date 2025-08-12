package main

import (
	"github.com/ropehapi/finance-manager-go/migrations"
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
	//TODO: Adicionar type à response de saida das transfers
	_ = godotenv.Load()

	database := db.NewDatabase()
	migrations.Migrate(database)

	accountRepo := repository.NewAccountRepository(database)
	transferRepo := repository.NewTransferRepository(database)
	paymentMethodRepo := repository.NewPaymentMethodRepository(database)
	debtRepo := repository.NewDebtRepository(database)

	accountService := service.NewAccountService(accountRepo)
	transferService := service.NewTransferService(transferRepo, accountRepo, paymentMethodRepo, debtRepo)
	paymentMethodService := service.NewPaymentMethodService(paymentMethodRepo)
	debtService := service.NewDebtService(debtRepo, accountRepo, transferRepo)

	accountHandler := handler.NewAccountHandler(accountService)
	transferHandler := handler.NewTransferHandler(transferService)
	paymentMethodHandler := handler.NewPaymentMethodHandler(paymentMethodService)
	debtHandler := handler.NewDebtHandler(debtService)

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

	debt := r.Group("/debts")
	debt.GET("/", debtHandler.GetAll)
	debt.POST("/pay/:id/:payer_account_id", debtHandler.Pay)
	debt.DELETE("/:id", debtHandler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r.Run(":" + port)

	log.Printf("Servidor rodando na porta %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
