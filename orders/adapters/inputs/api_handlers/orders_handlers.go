package api_handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/iamwy7/meli-challenge/orders/application/domain"
	"github.com/iamwy7/meli-challenge/orders/application/usecase"
	"github.com/iamwy7/meli-challenge/orders/application/usecase/dtos"
)

type OrdersHandler struct {
	createOrderUseCase *usecase.CreateOrderUseCase
	getOrderUseCase    *usecase.GetOrderUseCase
}

func NewOrderHandler(
	createOrderUseCase *usecase.CreateOrderUseCase,
	getOrderUseCase *usecase.GetOrderUseCase,
) *OrdersHandler {
	return &OrdersHandler{
		createOrderUseCase: createOrderUseCase,
		getOrderUseCase:    getOrderUseCase,
	}
}

func (h *OrdersHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input dtos.CreateOrderInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.writeErrorResponse(w, "invalid data to create an order", http.StatusBadRequest)
		return
	}
	output, err := h.createOrderUseCase.Execute(input)
	if err != nil {
		h.writeErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// @Param orderId path string true "Order ID"
func (h *OrdersHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	orderId := r.PathValue("orderId")
	output, err := h.getOrderUseCase.Execute(orderId)
	if errors.Is(err, domain.ErrOrderInvalidId) {
		h.writeErrorResponse(w, fmt.Sprintf("%v for orderId:%v", err.Error(), orderId), http.StatusNotFound)
		return
	}
	if err != nil {
		h.writeErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// writeErrorResponse writes an error response in JSON format
func (h *OrdersHandler) writeErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	log.Printf("handler error: %v", message)
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Message: message})
}

// ErrorResponse represents the structure of an error response
type ErrorResponse struct {
	Message string `json:"message"`
}
