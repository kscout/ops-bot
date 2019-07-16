package handlers

import (
	"github.com/gorilla/mux"
)

// Registerable can register itself with a router
type Registerable interface {
	// Register handler with router
	Register(router *mux.Router)
}
