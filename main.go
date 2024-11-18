package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Respond with "Hello, World!"
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	// Register the handler function to the root URL path
	http.HandleFunc("/", handler)

	// Start the server on port 8080
	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
