package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
)

// JSONResponder writes a JSON encoded object as a response
type JSONResponder struct {
	// Status wit h which to respond
	Status int
	
	// Data to write as JSON
	Data interface{}
}

// Respond with Status and Data JSON encoded
func (r JSONResponder) Respond(w http.ResponseWriter) {
	// Write Content-type header
	w.WriteHeader(r.Status)
	w.Header().Set("Content-type", "application/json")
	
	encoder := json.NewEncoder(w)
	if err := encoder.Encode(r.Data); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		
		_, wErr := fmt.Fprint(w, "{\"error\": \"internal server error\"}")
		if wErr != nil {
			panic(fmt.Errorf("error JSON encoding %#v: %s, then error "+
				"writing error response: %s", r.Data, err.Error(),
				wErr.Error()))
		}

		panic(fmt.Errorf("error JSON encoding %#v: %s", r.Data, err.Error()))
	}
}
