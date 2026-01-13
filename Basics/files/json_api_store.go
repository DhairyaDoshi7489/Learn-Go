package main

import (
	"encoding/json"
	"net/http"
)

// 1. Define the shape of your data
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

// 2. Create a "database" in memory using a Slice
var tasks = []Task{
	{ID: 1, Title: "Learn Go", Done: false},
	{ID: 2, Title: "Build an API", Done: false},
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	// Set the header so the client knows we are sending JSON
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		// Encode the 'tasks' slice into JSON and send it to the writer
		json.NewEncoder(w).Encode(tasks)

	case "POST":
		var newTask Task
		// Decode the JSON sent by the user into the 'newTask' struct
		// If the JSON is bad, return an error
		if err := json.NewDecoder(r.Body).Decode(&newTask); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Add to our "Database"
		tasks = append(tasks, newTask)

		// Return the created object to confirm
		json.NewEncoder(w).Encode(newTask)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/tasks", tasksHandler)
	http.ListenAndServe(":8080", nil)
}
