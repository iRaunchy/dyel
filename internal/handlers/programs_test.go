package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iraunchy/dyel/db"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockRepo struct {
	CreateFn func(ctx context.Context, p *db.Program) (*db.Program, error)
	GetFn    func(ctx context.Context, id string) (*db.Program, error)
	ListFn   func(ctx context.Context) ([]db.Program, error)
	UpdateFn func(ctx context.Context, p *db.Program) (*db.Program, error)
	DeleteFn func(ctx context.Context, id string) error
}

func (m *mockRepo) Create(ctx context.Context, p *db.Program) (*db.Program, error) {
	return m.CreateFn(ctx, p)
}
func (m *mockRepo) Get(ctx context.Context, id string) (*db.Program, error) {
	return m.GetFn(ctx, id)
}
func (m *mockRepo) List(ctx context.Context) ([]db.Program, error) { return m.ListFn(ctx) }
func (m *mockRepo) Update(ctx context.Context, p *db.Program) (*db.Program, error) {
	return m.UpdateFn(ctx, p)
}
func (m *mockRepo) Delete(ctx context.Context, id string) error { return m.DeleteFn(ctx, id) }

func TestCreateProgram_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	payload := db.Program{
		Name:     "Test Program",
		SharedBy: "test@example.com",
		Days:     []db.Day{}, // empty days for simplicity
	}
	returned := payload
	returned.ID = "uuid-123"

	repo := &mockRepo{
		CreateFn: func(ctx context.Context, p *db.Program) (*db.Program, error) {
			assert.NotNil(t, ctx)
			p.ID = returned.ID
			return p, nil
		},
	}
	h := NewHandler(repo)

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/programs", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.New()
	h.RegisterRoutes(router)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var got db.Program
	err := json.Unmarshal(w.Body.Bytes(), &got)
	assert.NoError(t, err)
	assert.Equal(t, returned.ID, got.ID)
	assert.Equal(t, payload.Name, got.Name)
	assert.Equal(t, payload.SharedBy, got.SharedBy)
}

func TestCreateProgram_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// fake that always returns an error
	repo := &mockRepo{
		CreateFn: func(ctx context.Context, p *db.Program) (*db.Program, error) {
			return nil, errors.New("db failure")
		},
	}
	h := NewHandler(repo)

	payload := db.Program{
		Name:     "Any Program",
		SharedBy: "bob@example.com",
		Days:     []db.Day{},
	}
	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/programs", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router := gin.New()
	h.RegisterRoutes(router)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var errResp map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.Contains(t, errResp["error"], "db failure")
}

func TestGetProgram_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expected := &db.Program{
		ID:       "uuid-123",
		Name:     "Test Program",
		SharedBy: "test@example.com",
		Days:     []db.Day{},
	}

	repo := &mockRepo{
		GetFn: func(ctx context.Context, id string) (*db.Program, error) {
			assert.Equal(t, "uuid-123", id)
			return expected, nil
		},
	}
	h := NewHandler(repo)
	router := gin.New()
	h.RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/programs/"+expected.ID, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actual db.Program
	err := json.Unmarshal(w.Body.Bytes(), &actual)
	assert.NoError(t, err)
	assert.Equal(t, expected.ID, actual.ID)
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.SharedBy, actual.SharedBy)
}

func TestListProgram_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	expectedPrograms := []db.Program{
		{
			ID:       "uuid-123",
			Name:     "Test Program 1",
			SharedBy: "test@example.com",
			Days:     []db.Day{},
		},
		{
			ID:       "uuid-456",
			Name:     "Test Program 2",
			SharedBy: "test2@example.com",
			Days:     []db.Day{},
		},
	}

	repo := &mockRepo{
		ListFn: func(ctx context.Context) ([]db.Program, error) {
			assert.NotNil(t, ctx)
			return expectedPrograms, nil
		},
	}

	h := NewHandler(repo)
	router := gin.New()
	h.RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/programs", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var actualPrograms []db.Program
	err := json.Unmarshal(w.Body.Bytes(), &actualPrograms)
	assert.NoError(t, err)
	assert.Equal(t, expectedPrograms, actualPrograms)
}

func TestUpdateProgram_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	validID := "123e4567-e89b-12d3-a456-426614174000"
	expected := &db.Program{
		ID:       validID,
		Name:     "Updated Name",
		SharedBy: "bob@example.com",
		Days:     []db.Day{},
	}

	repo := &mockRepo{
		UpdateFn: func(ctx context.Context, p *db.Program) (*db.Program, error) {
			assert.Equal(t, expected, p)
			return expected, nil
		},
	}

	h := NewHandler(repo)
	router := gin.New()
	h.RegisterRoutes(router)

	payload := UpdateProgramJSON{
		Name:     expected.Name,
		SharedBy: expected.SharedBy,
		Days:     expected.Days,
	}
	bodyBytes, err := json.Marshal(payload)
	assert.NoError(t, err)

	req := httptest.NewRequest(
		http.MethodPut,
		"/api/v1/programs/"+validID,
		bytes.NewReader(bodyBytes),
	)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var actual db.Program
	assert.NoError(t, json.Unmarshal(w.Body.Bytes(), &actual))
	assert.Equal(t, expected, &actual)
}

func TestUpdateProgram_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	bodyBytes, _ := json.Marshal(map[string]interface{}{
		"name":      "X",
		"shared_by": "bob@example.com",
		"days":      []db.Day{},
	})

	// wrap Update to always fail
	repo := &mockRepo{
		UpdateFn: func(ctx context.Context, p *db.Program) (*db.Program, error) {
			return nil, errors.New("update failed")
		},
	}

	h := NewHandler(repo)
	router := gin.New()
	h.RegisterRoutes(router)

	validID := "00000000-0000-0000-0000-000000000000"
	url := fmt.Sprintf("/api/v1/programs/%s", validID)
	req := httptest.NewRequest(http.MethodPut, url, bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)

	var errResp map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.Contains(t, errResp["error"], "update failed")
}

func TestDeleteProgram_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	called := false
	repo := &mockRepo{
		DeleteFn: func(ctx context.Context, id string) error {
			called = true
			assert.Equal(t, "uuid-789", id)
			assert.NotNil(t, ctx)
			return nil
		},
	}

	h := NewHandler(repo)
	router := gin.New()
	h.RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/programs/uuid-789", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.True(t, called, "DeleteFn should have been called")
	assert.Equal(t, http.StatusNoContent, w.Code)
	assert.Empty(t, w.Body.String())
}

func TestDeleteProgram_Error(t *testing.T) {
	gin.SetMode(gin.TestMode)

	repo := &mockRepo{
		DeleteFn: func(ctx context.Context, id string) error {
			return errors.New("delete failed")
		},
	}

	h := NewHandler(repo)
	router := gin.New()
	h.RegisterRoutes(router)

	req := httptest.NewRequest(http.MethodDelete, "/api/v1/programs/any-id", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var errResp map[string]string
	_ = json.Unmarshal(w.Body.Bytes(), &errResp)
	assert.Contains(t, errResp["error"], "delete failed")
}
