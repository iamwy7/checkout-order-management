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

// OrdersHandler is our principal api handler
type OrdersHandler struct {
	createOrderUseCase *usecase.CreateOrderUseCase
	getOrderUseCase    *usecase.GetOrderUseCase
}

// NewOrderHandler creates an instance of OrdersHandler
func NewOrderHandler(
	createOrderUseCase *usecase.CreateOrderUseCase,
	getOrderUseCase *usecase.GetOrderUseCase,
) *OrdersHandler {
	return &OrdersHandler{
		createOrderUseCase: createOrderUseCase,
		getOrderUseCase:    getOrderUseCase,
	}
}

// CreateOrder handles the request to create an order
//
//	@Summary		Create an order
//	@Description	Create an order with specific products
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			input	body		dtos.CreateOrderInputDTO	true	"Order Input"
//	@Success		201		{object}	dtos.CreatedOrderOutputDTO
//	@Failure		400		{object}	ErrorResponse
//	@Failure		404		{object}	ErrorResponse
//	@Failure		500		{object}	ErrorResponse
//	@Router			/orders [post]
func (h *OrdersHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var input dtos.CreateOrderInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		h.writeErrorResponse(w, "invalid data to create an order", http.StatusBadRequest)
		return
	}
	output, err := h.createOrderUseCase.Execute(input)
	if errors.Is(err, domain.ErrOrderWithTooMuchProducts) {
		h.writeErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err != nil {
		h.writeErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(output)
}

// GetOrder handles the request to get an existing order
//
//	@Summary		List all events
//	@Description	Get all events with their details
//	@Tags			Orders
//	@Accept			json
//	@Produce		json
//	@Param			orderId	path		string	true	"orderId"
//	@Success		200		{object}	dtos.CreatedOrderOutputDTO
//	@Failue			422 {object} ErrorResponse
//	@Failure		500	{object}	ErrorResponse
//	@Router			/orders/{orderId} [get]
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
	Message string `json:"message" example:" message about the error"`
}
