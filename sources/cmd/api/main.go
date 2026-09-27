package main

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
