package main

import (
	"fmt"
	"net/http"
)

// This function handles requests to the root path "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	// w is where you write the response
	// r contains details about the request (method, headers, body)
	fmt.Fprintf(w, "Welcome to the Home Page!")
}

// This handles requests to "/about"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a simple Go web server.")
}

// This handles requests to "/register"
func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You have successfully registered.")
}
func main() {
	// Register the routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("Server starting on port 8080...")

	// Start the server (this blocks forever)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
