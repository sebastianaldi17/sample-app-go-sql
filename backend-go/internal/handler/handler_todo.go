package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	todoEntity "github.com/sebastianaldi17/sample-app-go-sql/internal/entity/todo"
)

func (h *Handler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not a number"))
		return
	}

	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID in JWT"))
		return
	}

	todo, err := h.uc.GetTodoByID(todoIDInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if todo.AuthorID != userID {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("User has no access to this todo"))
		return
	}

	todoBytes, err := json.Marshal(todo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(todoBytes)
}

func (h *Handler) InsertTodo(w http.ResponseWriter, r *http.Request) {
	var req todoEntity.InsertTodoRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID in JWT"))
		return
	}

	req.UserID = userID

	if err := h.uc.InsertTodo(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}

func (h *Handler) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Todo ID is not a number"))
		return
	}

	var req todoEntity.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID in JWT"))
		return
	}

	isAuthor, err := h.uc.VerifyTodoAuthor(todoIDInt, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !isAuthor {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Todo author is not the same as current user"))
		return
	}

	req.ID = todoIDInt

	if err = h.uc.UpdateTodo(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}

func (h *Handler) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not a number"))
		return
	}

	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID in JWT"))
		return
	}

	isAuthor, err := h.uc.VerifyTodoAuthor(todoIDInt, userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	if !isAuthor {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Todo author is not the same as current user"))
		return
	}

	if err = h.uc.DeleteTodo(todoIDInt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}

func (h *Handler) GetTodosByUser(w http.ResponseWriter, r *http.Request) {
	userID, err := getUserIDFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid user ID in JWT"))
		return
	}

	todos, err := h.uc.GetTodoByAuthor(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	todoBytes, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(todoBytes)
}
