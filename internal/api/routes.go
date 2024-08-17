package api

import "github.com/gorilla/mux"

func SetupRoutes(h *TransactionHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/transactionservice/transaction/{id}", h.CreateTransaction).Methods("PUT")
	r.HandleFunc("/transactionservice/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transactionservice/transaction/types/{type}", h.GetTransactionsByType).Methods("GET")
	r.HandleFunc("/transactionservice/transaction/sum{id}", h.GetTransactionSum).Methods("GET")

	return r
}
