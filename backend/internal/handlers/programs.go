package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iraunchy/dyel/backend/db"
)

func (h *Handler) CreateProgram(c *gin.Context) {
	HandleJSON[CreateProgramInput, *db.Program](
		c,
		BindJSON[CreateProgramInput],
		func(ctx context.Context, in CreateProgramInput) (*db.Program, error) {
			return h.Repo.Create(ctx, in.ToModel())
		},
		http.StatusCreated,
	)
}

// ListPrograms handles GET /api/v1/programs
func (h *Handler) ListPrograms(c *gin.Context) {
	HandleJSON[ListProgramsInput, []db.Program](
		c,
		func(c *gin.Context) (ListProgramsInput, error) {
			return ListProgramsInput{}, nil
		},
		func(ctx context.Context, _ ListProgramsInput) ([]db.Program, error) {
			return h.Repo.List(ctx)
		},
		http.StatusOK,
	)
}

// GetProgram handles GET /api/v1/programs/:id
func (h *Handler) GetProgram(c *gin.Context) {
	HandleJSON[GetProgramInput, *db.Program](
		c,
		BindURI[GetProgramInput],
		func(ctx context.Context, in GetProgramInput) (*db.Program, error) {
			return h.Repo.Get(ctx, in.ID)
		},
		http.StatusOK,
	)
}

// UpdateProgram handles PUT /api/v1/programs/:id
func (h *Handler) UpdateProgram(c *gin.Context) {
	HandleJSON[UpdateProgramInput, *db.Program](
		c,
		func(c *gin.Context) (UpdateProgramInput, error) {
			uri, err := BindURI[UpdateProgramURI](c)
			if err != nil {
				return UpdateProgramInput{}, err
			}
			body, err := BindJSON[UpdateProgramJSON](c)
			return uri.Merge(body), err
		},
		func(ctx context.Context, in UpdateProgramInput) (*db.Program, error) {
			return h.Repo.Update(ctx, in.ToModel())
		},
		http.StatusOK,
	)
}

// DeleteProgram handles DELETE /api/v1/programs/:id
func (h *Handler) DeleteProgram(c *gin.Context) {
	HandleJSON[string, struct{}](
		c,
		func(c *gin.Context) (string, error) {
			return c.Param("id"), nil
		},
		func(ctx context.Context, id string) (struct{}, error) {
			if err := h.Repo.Delete(ctx, id); err != nil {
				return struct{}{}, err
			}
			return struct{}{}, nil
		},
		http.StatusNoContent,
	)
}
