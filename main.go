package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler function for the root endpoint
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Hello, World!")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Register the handler function for the root endpoint
	r.HandleFunc("/", helloHandler).Methods(http.MethodGet)

	// Start the server on port 3001
	fmt.Println("Server is listening on port 3001...")
	if err := http.ListenAndServe(":3001", r); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
