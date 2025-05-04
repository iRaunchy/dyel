package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iraunchy/dyel/db"
	httpresp "github.com/iraunchy/dyel/internal/http"
)

// createProgramInput holds the JSON‐binding rules for create & update.
type createProgramInput struct {
	Name     string   `json:"name" binding:"required"`
	SharedBy string   `json:"shared_by" binding:"required"`
	Days     []db.Day `json:"days" binding:"required"`
}

func (in *createProgramInput) toDB() *db.Program {
	return &db.Program{
		Name:     in.Name,
		SharedBy: in.SharedBy,
		Days:     in.Days,
	}
}

// CreateProgram handles POST /api/v1/programs
func (h *Handler) CreateProgram(c *gin.Context) {
	var in createProgramInput
	if err := c.ShouldBindJSON(&in); err != nil {
		httpresp.Error(c, http.StatusBadRequest, err)
		return
	}

	prog, err := h.Repo.Create(c.Request.Context(), in.toDB())
	if err != nil {
		httpresp.Error(c, http.StatusInternalServerError, err)
		return
	}

	httpresp.Created(c, prog)
}

// ListPrograms handles GET /api/v1/programs
func (h *Handler) ListPrograms(c *gin.Context) {
	list, err := h.Repo.List(c.Request.Context())
	if err != nil {
		httpresp.Error(c, http.StatusInternalServerError, err)
		return
	}
	httpresp.JSON(c, http.StatusOK, list)
}

// GetProgram handles GET /api/v1/programs/:id
func (h *Handler) GetProgram(c *gin.Context) {
	id := c.Param("id")
	prog, err := h.Repo.Get(c.Request.Context(), id)
	if err != nil {
		httpresp.Error(c, http.StatusInternalServerError, err)
		return
	}
	httpresp.JSON(c, http.StatusOK, prog)
}

// UpdateProgram handles PUT /api/v1/programs/:id
func (h *Handler) UpdateProgram(c *gin.Context) {
	var in createProgramInput
	if err := c.ShouldBindJSON(&in); err != nil {
		httpresp.Error(c, http.StatusBadRequest, err)
		return
	}

	id := c.Param("id")
	upd := in.toDB()
	upd.ID = id

	prog, err := h.Repo.Update(c.Request.Context(), upd)
	if err != nil {
		httpresp.Error(c, http.StatusInternalServerError, err)
		return
	}
	httpresp.JSON(c, http.StatusOK, prog)
}

// DeleteProgram handles DELETE /api/v1/programs/:id
func (h *Handler) DeleteProgram(c *gin.Context) {
	id := c.Param("id")
	if err := h.Repo.Delete(c.Request.Context(), id); err != nil {
		httpresp.Error(c, http.StatusInternalServerError, err)
		return
	}
	c.Status(http.StatusNoContent)
}
