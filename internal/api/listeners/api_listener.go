package listeners

import (
	"fmt"
	"net/http"
	"schoolapi/internal/api/handlers"
)

func StartAPIServer(port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HandleRoot)

	connectionString := fmt.Sprintf(":%d", port)
	fmt.Println(fmt.Sprintf("API server listening on %s", connectionString))
	if err := http.ListenAndServe(connectionString, mux); err != nil {
		fmt.Println("Error starting server:", err)
	}
}