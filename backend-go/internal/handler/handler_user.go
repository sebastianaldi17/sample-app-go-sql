package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	userEntity "github.com/sebastianaldi17/sample-app-go-sql/internal/entity/user"
)

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var req userEntity.Login

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := h.uc.CreateAccount(req); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("OK"))
}

func (h *Handler) LoginUser(w http.ResponseWriter, r *http.Request) {
	var req userEntity.Login

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	if err := h.uc.ValidateLogin(req); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(err.Error()))
		return
	}

	token, err := h.uc.CreateJWT(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"token":"%s"}`, token)))
}

func (h *Handler) ValidateJWT(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
