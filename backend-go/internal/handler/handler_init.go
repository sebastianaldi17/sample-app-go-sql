package handler

import (
	"net/http"

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
