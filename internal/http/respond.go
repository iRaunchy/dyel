package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// JSON writes any object at the given status code.
func JSON(c *gin.Context, code int, obj interface{}) {
	c.JSON(code, obj)
}

// Error writes a JSON error body with {"error": "..."}.
func Error(c *gin.Context, code int, err error) {
	c.JSON(code, gin.H{"error": err.Error()})
}

// Created writes a 201 with the given object.
func Created(c *gin.Context, obj interface{}) {
	JSON(c, http.StatusCreated, obj)
}
