package main

import (
	"finance-api/config"
	"finance-api/internal/handler"
	"finance-api/internal/middleware"
	"finance-api/internal/repository"
	"finance-api/internal/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	db := config.ConnectDB()
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	transactionRepo := repository.NewTransactionRepository(db)

	authService := service.NewAuthService(userRepo)
	transactionService := service.NewTransactionService(transactionRepo)

	authHandler := handler.NewAuthHandler(authService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	r := chi.NewRouter()

	// rotas públicas
	r.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.Register)
		r.Post("/login", authHandler.Login)
	})

	// rotas protegidas
	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)
		r.Post("/transactions", transactionHandler.CreateTransaction)
		r.Get("/transactions", transactionHandler.GetTransactions)
		r.Delete("/transactions/{id}", transactionHandler.DeleteTransaction)
		r.Get("/exchange", handler.GetExchangeRate)
	})
	log.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", r)
}
