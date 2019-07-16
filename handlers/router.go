package handlers

import (
	"net/http"
	
	"github.com/gorilla/mux"
)

// Router defines routes via Registerables
type Router struct {
	// router is the backend which routes and runs the correct handler
	// based on parameters provided in Registerables
	router *mux.Router
}

// NewRouter initializes a Router
func NewRouter() *Router {
	return &Router{
		router: mux.NewRouter(),
	}
}

// Add Registerable to router
func (r *Router) Add(registerable Registerable) {
	registerable.Register(r.router)
}

// ServeHTTP serve the handlers registered by Registerables
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
