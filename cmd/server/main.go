package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/avalokitasharma/transaction_service/transaction_service/internal/api"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/repository"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/service"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.NewTransactionoRepository(db)
	service := service.NewTransactionService(repo)
	handler := api.NewTransactionoHandler(service)

	router := api.SetupRoutes(handler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
