package handlers

import (
	"github.com/iraunchy/dyel/backend/internal/repos"
)

type Handler struct {
	Repo repos.ProgramRepo
}

// NewHandler wires in a ProgramRepo
func NewHandler(r repos.ProgramRepo) *Handler {
	return &Handler{Repo: r}
}
