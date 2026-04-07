package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var idCounter int
var tasks []Task

func main() {

	http.HandleFunc("/tasks", addHandler)         // post
	http.HandleFunc("/tasks/{id}", updateHandler) // put

	port := ":8080"
	log.Printf("Starting Server localhost %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Created_At  time.Time `json:"created_at"`
	Updated_At  time.Time `json:"updated_at"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	var req CreateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}
	if req.Description == "" {
		writeError(w, http.StatusBadRequest, "write something in description")
		return
	}

	idCounter++

	newTask := Task{
		ID:          idCounter,
		Description: req.Description,
		Status:      "in progress",
		Created_At:  time.Now(),
		Updated_At:  time.Now(),
	}

	tasks = append(tasks, newTask)
	writeJSON(w, http.StatusOK, APIResponse{
		Success: true,
		Data:    newTask,
	})
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	var req UpdateTaskRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if req.Description == nil && req.Status == nil {
		writeError(w, http.StatusBadRequest, "nothing to update")
		return
	}

	for i := range tasks {
		if tasks[i].ID == req.ID {
			if req.Description != nil {
				tasks[i].Description = *req.Description
			}
			if req.Status != nil {
				tasks[i].Status = *req.Status
			}
			tasks[i].Updated_At = time.Now()
			writeJSON(w, http.StatusOK, APIResponse{
				Success: true,
				Data:    tasks[i],
			})
			return
		}
	}
	writeError(w, http.StatusNotFound, "task not found")
}

type CreateTaskRequest struct {
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	ID          int       `json:"id"`
	Description *string   `json:"description"`
	Status      *string   `json:"status"`
	Updated_At  time.Time `json:"updated_at"`
}

type APIResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, APIResponse{
		Success: false,
		Error:   message,
	})
}
