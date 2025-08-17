package handlers

import (
	"fmt"
	"net/http"
)

// HandleRoot responds to requests at the root endpoint.
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}