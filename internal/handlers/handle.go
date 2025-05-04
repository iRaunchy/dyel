package handlers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	httpresp "github.com/iraunchy/dyel/internal/http"
)

// BinderFn is any function that reads from gin.Context and returns
// a typed input or an error.
type BinderFn[T any] func(*gin.Context) (T, error)

// ActionFn is any function that, given a context and a typed input,
// returns a typed output or an error.
type ActionFn[T any, R any] func(context.Context, T) (R, error)

func HandleJSON[T any, R any](
	c *gin.Context,
	binder BinderFn[T],
	action ActionFn[T, R],
	successCode int,
) {
	in, err := binder(c)
	if err != nil {
		fmt.Printf("binder error = %v\n", err)
		httpresp.Error(c, http.StatusBadRequest, err)
		return
	}

	out, err := action(c.Request.Context(), in)
	if err != nil {
		httpresp.Error(c, http.StatusInternalServerError, err)
		return
	}

	if successCode == http.StatusCreated {
		httpresp.Created(c, out)
	} else {
		httpresp.JSON(c, successCode, out)
	}
}

func BindURI[T any](c *gin.Context) (T, error) {
	var in T
	return in, c.ShouldBindUri(&in)
}
func BindJSON[T any](c *gin.Context) (T, error) {
	var in T
	return in, c.ShouldBindJSON(&in)
}
