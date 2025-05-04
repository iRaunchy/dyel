package handlers

import "github.com/gin-gonic/gin"

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		api.POST("/programs", h.CreateProgram)
		api.GET("/programs", h.ListPrograms)
		api.GET("/programs/:id", h.GetProgram)
		api.PUT("/programs/:id", h.UpdateProgram)
		api.DELETE("/programs/:id", h.DeleteProgram)
	}
}
