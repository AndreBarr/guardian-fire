package main

import (
	"fmt"
	"net/http"
)

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func fetchMessage(w http.ResponseWriter, r *http.Request) {
	// Simulate fetching data
	message := "This is a dynamic message fetched from the server!"
	fmt.Fprint(w, message)
}

func main() {
	// Serve static files (e.g., HTMX script)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register the handler function to the root URL path
	http.HandleFunc("/", serveHTML)

	// Handle HTMX requests
	http.HandleFunc("/fetch-message", fetchMessage)

	// Start the server on port 8080
	fmt.Println("Starting server at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
