package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	files := map[string]string{
		"go.mod": `module go-clean-arch

go 1.22
`,
		"internal/domain/user.go": `package domain

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("usuario no encontrado")

type User struct {
	ID    int64  ` + "`json:\"id\"`" + `
	Name  string ` + "`json:\"name\"`" + `
	Email string ` + "`json:\"email\"`" + `
}

type UserRepository interface {
	GetByID(ctx context.Context, id int64) (*User, error)
}

type UserUseCase interface {
	GetUser(ctx context.Context, id int64) (*User, error)
}
`,
		"internal/usecase/user_usecase.go": `package usecase

import (
	"context"
	"go-clean-arch/internal/domain"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repo: r}
}

func (u *userUseCase) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return u.repo.GetByID(ctx, id)
}
`,
		"internal/delivery/http/user_handler.go": `package http

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
`,
		"cmd/api/main.go": `package main

import (
	"context"
	"log"
	"net/http"

	delivery "go-clean-arch/internal/delivery/http"
	"go-clean-arch/internal/domain"
	"go-clean-arch/internal/usecase"
)

// inMemoryRepo simula la base de datos
type inMemoryRepo struct{}

func (m *inMemoryRepo) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	if id == 1 {
		return &domain.User{ID: 1, Name: "Gopher", Email: "gopher@go.dev"}, nil
	}
	return nil, domain.ErrUserNotFound
}

func main() {
	repo := &inMemoryRepo{}
	userUC := usecase.NewUserUseCase(repo)
	userHandler := delivery.NewUserHandler(userUC)

	mux := http.NewServeMux()
	userHandler.RegisterRoutes(mux)

	log.Println("⚡ Servidor escuchando en http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}
`,
	}

	for path, content := range files {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			fmt.Println("❌ Error creando directorio:", err)
			return
		}
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println("❌ Error creando archivo:", err)
			return
		}
	}
	
	fmt.Println("✅ ¡Estructura base de Clean Architecture generada con éxito!")
	fmt.Println("🚀 Para arrancar la API ejecuta: go run cmd/api/main.go")
}