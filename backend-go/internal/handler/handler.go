package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/entity"
	"github.com/sebastianaldi17/sample-app-go-sql/internal/usecase"
)

type Handler struct {
	uc usecase.Usecase
}

func New(usecase usecase.Usecase) *Handler {
	return &Handler{
		uc: usecase,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World!"))
}

func (h *Handler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.uc.GetTodos()
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

func (h *Handler) GetTodoByID(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	todoIDInt, err := strconv.ParseInt(todoID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("ID is not a number"))
		return
	}
	todo, err := h.uc.GetTodoByID(todoIDInt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
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
	var req entity.InsertTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
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
		w.Write([]byte("ID is not a number"))
		return
	}

	var req entity.UpdateTodoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
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
	if err = h.uc.DeleteTodo(todoIDInt); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}
