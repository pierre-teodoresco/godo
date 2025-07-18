package server

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"
	"todo-web-go-vue/internal/repository"

	"github.com/google/uuid"
)

// Custom error type
type ApiError struct {
	Err        error
	StatusCode int
}

func (e *ApiError) Error() string {
	return e.Err.Error()
}

// Parsing type
type CreateTaskRequest struct {
	Title string `json:"title"`
}

type DeleteTaskRequest struct {
	ID uuid.UUID `json:"id"`
}

type ApiTask struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt *string   `json:"created_at"`
}

func unmarshallJsonRequest[T CreateTaskRequest | DeleteTaskRequest | repository.UpdateTaskParams](r *http.Request) (T, *ApiError) {
	// Default T value
	var req T

	// Check content-type looking for application/json
	if r.Header.Get("Content-Type") != "" && !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
		return req, &ApiError{
			Err:        errors.New("content-Type must be application/json"),
			StatusCode: http.StatusBadRequest,
		}
	}

	// Unmarshall JSON from request into req
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&req)
	if err != nil {
		// Check if the error is due to invalid JSON structure
		if _, ok := err.(*json.SyntaxError); ok {
			return req, &ApiError{
				Err:        errors.New("invalid JSON format: malformed JSON structure"),
				StatusCode: http.StatusBadRequest,
			}
		}
		// Check if the error is due to an invalid parameter type
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return req, &ApiError{
				Err:        errors.New("invalid field type"),
				StatusCode: http.StatusBadRequest,
			}
		}
		// Check for unknown fields
		if strings.Contains(err.Error(), "unknown field") {
			return req, &ApiError{
				Err:        errors.New("invalid request format"),
				StatusCode: http.StatusBadRequest,
			}
		}
		return req, &ApiError{
			Err:        errors.New("failed to unmarshall request"),
			StatusCode: http.StatusInternalServerError,
		}
	}
	return req, nil
}

func taskToApiTask(t repository.Task) ApiTask {
	var createdAt *string
	if t.CreatedAt.Valid {
		s := t.CreatedAt.Time.Format(time.RFC3339)
		createdAt = &s
	}
	return ApiTask{
		ID:        t.ID,
		Title:     t.Title,
		Completed: t.Completed,
		CreatedAt: createdAt,
	}
}

func sendJsonTaskResponse[T ApiTask | []ApiTask](w http.ResponseWriter, data T, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to marshal create task response", http.StatusInternalServerError)
		return
	}
}

/*
================
	HANDLERS
================
*/

func (s *Server) ListAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := s.db.GetQueries().FindAllTasks(context.Background())
	if err != nil {
		http.Error(w, "Failed to select all tasks", http.StatusInternalServerError)
		return
	}

	// Convert tasks to []ApiTask[]
	var response []ApiTask
	for _, task := range tasks {
		response = append(response, taskToApiTask(task))
	}

	sendJsonTaskResponse(w, response, http.StatusOK)
}

func (s *Server) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {
	req, e := unmarshallJsonRequest[CreateTaskRequest](r)
	if e != nil {
		http.Error(w, e.Error(), e.StatusCode)
		return
	}

	task, err := s.db.GetQueries().CreateTask(context.Background(), req.Title)
	if err != nil {
		http.Error(w, "Failed to created new task", http.StatusInternalServerError)
		return
	}

	// Convert task to ApiTask
	response := taskToApiTask(task)

	sendJsonTaskResponse(w, response, http.StatusCreated)
}

func (s *Server) DeleteTask(w http.ResponseWriter, r *http.Request) {
	req, e := unmarshallJsonRequest[DeleteTaskRequest](r)
	if e != nil {
		http.Error(w, e.Error(), e.StatusCode)
		return
	}

	err := s.db.GetQueries().DeleteTask(context.Background(), req.ID)
	if err != nil {
		http.Error(w, "Failed to delete task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (s *Server) UpdateTask(w http.ResponseWriter, r *http.Request) {
	req, e := unmarshallJsonRequest[repository.UpdateTaskParams](r)
	if e != nil {
		http.Error(w, e.Error(), e.StatusCode)
		return
	}

	task, err := s.db.GetQueries().UpdateTask(context.Background(), req)
	if err != nil {
		http.Error(w, "Failed to update task", http.StatusInternalServerError)
		return
	}

	// Convert task to ApiTask
	response := taskToApiTask(task)

	sendJsonTaskResponse(w, response, http.StatusOK)
}
