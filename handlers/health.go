package handlers

import (
	"net/http"
	
	"github.com/gorilla/mux"
)

// HealthHandler returns if the server is functioning correctly
type HealthHandler struct{}

// HealthResponse holds status of server
type HealthResponse struct {
	// OK indicates if server is working correctly
	OK bool `json:"ok"`
}

// Register handler at GET /health
func (h HealthHandler) Register(router *mux.Router) {
	router.Handle("/health", h).Methods("GET")
}

// ServeHTTP responds to all requests with a HealthResponse and status HTTP 200 OK
func (h HealthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := JSONResponder{
		Status: http.StatusOK,
		Data: HealthResponse{
			OK: true,
		},
	}
	resp.Respond(w)
}
