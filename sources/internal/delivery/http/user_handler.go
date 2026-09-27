package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"go-clean-arch/internal/domain"
)

type UserHandler struct {
	useCase domain.UserUseCase
}

func NewUserHandler(uc domain.UserUseCase) *UserHandler {
	return &UserHandler{useCase: uc}
}

func (h *UserHandler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users/", h.handleUserByID)
}

func (h *UserHandler) handleUserByID(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "ID no provisto", http.StatusBadRequest)
		return
	}

	id, _ := strconv.ParseInt(parts[1], 10, 64)
	user, err := h.useCase.GetUser(r.Context(), id)
	
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
