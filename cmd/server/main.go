package main

import (
	"database/sql"
	"log"
	"net/http"
	"path/filepath"

	"github.com/avalokitasharma/transaction_service/transaction_service/internal/api"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/repository"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/service"
	"github.com/avalokitasharma/transaction_service/transaction_service/pkg/database"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://username:password@localhost/dbname?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	migrationsPath, err := filepath.Abs("./migrations")
	if err != nil {
		log.Fatal(err)
	}
	err = database.RunMigrations(db, migrationsPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.NewTransactionoRepository(db)
	service := service.NewTransactionService(repo)
	handler := api.NewTransactionHandler(service)

	router := api.SetupRoutes(handler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
