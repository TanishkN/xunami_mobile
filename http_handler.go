package http_handler

import (
	"fmt"
	"net/http"
)

type handler struct {
	//gateway

}

func NewHandler() *handler {
	return &handler{}

}

// RegisterRoutes registers all the routes with the mux
func (h *handler) registerRoutes(mux *http.ServeMux) {
	// Example route: POST /api/create-order
	mux.HandleFunc("/api/create-order", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	// Ensure this is a POST request
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Your logic to handle order creation goes here
	fmt.Fprintf(w, "Order created successfully")
}
