package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/avalokitasharma/transaction_service/transaction_service/internal/models"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/service"
	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	service *service.TransactionService
}

func NewTransactionHandler(s *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{service: s}
}

func (h *TransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid transaction id", http.StatusBadRequest)
		return
	}
	var t models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.ID = id
	if err := h.service.CreateTransaction(&t); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *TransactionHandler) GetTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid transaction id", http.StatusBadRequest)
		return
	}
	t, err := h.service.GetTransaction(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(t)
}
func (h *TransactionHandler) GetTransactionsByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionType := vars["type"]
	ids, err := h.service.GetTransactionsByType(transactionType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(ids)
}

func (h *TransactionHandler) GetTransactionSum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid transaction id", http.StatusBadRequest)
		return
	}
	sum, err := h.service.GetTransactionSum(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]float64{"sum": sum})
}
