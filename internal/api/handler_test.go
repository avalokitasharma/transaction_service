package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/avalokitasharma/transaction_service/transaction_service/internal/api"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/models"
	"github.com/avalokitasharma/transaction_service/transaction_service/internal/service"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type MockService struct {
	CreateTransactionFunc     func(*models.Transaction) error
	GetTransactionFunc        func(int64) (*models.Transaction, error)
	GetTransactionsByTypeFunc func(string) ([]int64, error)
	GetTransactionSumFunc     func(int64) (float64, error)
}

func (m *MockService) CreateTransaction(t *models.Transaction) error {
	return m.CreateTransactionFunc(t)
}

func (m *MockService) GetTransaction(id int64) (*models.Transaction, error) {
	return m.GetTransactionFunc(id)
}

func (m *MockService) GetTransactionsByType(transactionType string) ([]int64, error) {
	return m.GetTransactionsByTypeFunc(transactionType)
}

func (m *MockService) GetTransactionSum(id int64) (float64, error) {
	return m.GetTransactionSumFunc(id)
}

//var _ service.TransactionService = (*MockService)(nil)

var _ = Describe("TransactionHandler", func() {
	var (
		mockService *MockService
		handler     *api.TransactionHandler
		router      *mux.Router
	)

	BeforeEach(func() {
		mockService = &MockService{}
		handler = api.NewTransactionHandler(mockService)
		router = mux.NewRouter()
		router.HandleFunc("/transactionservice/transaction/{id}", handler.CreateTransaction).Methods("PUT")
		router.HandleFunc("/transactionservice/transaction/{id}", handler.GetTransaction).Methods("GET")
		router.HandleFunc("/transactionservice/types/{type}", handler.GetTransactionsByType).Methods("GET")
		router.HandleFunc("/transactionservice/sum/{id}", handler.GetTransactionSum).Methods("GET")
	})

	Describe("CreateTransaction", func() {
		It("should create a transaction successfully", func() {
			transaction := &models.Transaction{
				ID:     10,
				Amount: 5000,
				Type:   "cars",
			}

			mockService.CreateTransactionFunc = func(t *models.Transaction) error {
				Expect(t).To(Equal(transaction))
				return nil
			}

			body, _ := json.Marshal(transaction)
			req, _ := http.NewRequest("PUT", "/transactionservice/transaction/10", bytes.NewBuffer(body))
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))

			var response map[string]string
			json.Unmarshal(rr.Body.Bytes(), &response)
			Expect(response["status"]).To(Equal("ok"))
		})
	})

	Describe("GetTransaction", func() {
		It("should get a transaction successfully", func() {
			transaction := &models.Transaction{
				ID:     10,
				Amount: 5000,
				Type:   "cars",
			}

			mockService.GetTransactionFunc = func(id int64) (*models.Transaction, error) {
				Expect(id).To(Equal(int64(10)))
				return transaction, nil
			}

			req, _ := http.NewRequest("GET", "/transactionservice/transaction/10", nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))

			var response models.Transaction
			json.Unmarshal(rr.Body.Bytes(), &response)
			Expect(response).To(Equal(*transaction))
		})
	})

	Describe("GetTransactionsByType", func() {
		It("should get transactions by type successfully", func() {
			transactionIDs := []int64{1, 2, 3}

			mockService.GetTransactionsByTypeFunc = func(transactionType string) ([]int64, error) {
				Expect(transactionType).To(Equal("cars"))
				return transactionIDs, nil
			}

			req, _ := http.NewRequest("GET", "/transactionservice/types/cars", nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))

			var response []int64
			json.Unmarshal(rr.Body.Bytes(), &response)
			Expect(response).To(Equal(transactionIDs))
		})
	})

	Describe("GetTransactionSum", func() {
		It("should get transaction sum successfully", func() {
			sum := 15000.0

			mockService.GetTransactionSumFunc = func(id int64) (float64, error) {
				Expect(id).To(Equal(int64(10)))
				return sum, nil
			}

			req, _ := http.NewRequest("GET", "/transactionservice/sum/10", nil)
			rr := httptest.NewRecorder()

			router.ServeHTTP(rr, req)

			Expect(rr.Code).To(Equal(http.StatusOK))

			var response map[string]float64
			json.Unmarshal(rr.Body.Bytes(), &response)
			Expect(response["sum"]).To(Equal(sum))
		})
	})
})
`