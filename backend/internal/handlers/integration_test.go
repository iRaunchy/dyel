package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/iraunchy/dyel/backend/db"
	"github.com/iraunchy/dyel/backend/internal/repos"
)

func setupRouter(t *testing.T) *gin.Engine {
	dbConn, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	assert.NoError(t, err)

	assert.NoError(t, dbConn.AutoMigrate(&db.Program{}, &db.Day{}, &db.Exercise{}))

	repo := repos.NewGORMProgramRepo(dbConn)

	h := NewHandler(repo)

	r := gin.New()
	r.Use(gin.Recovery())
	h.RegisterRoutes(r)

	return r
}

func TestCreateAndFetchProgram(t *testing.T) {
	router := setupRouter(t)

	payload := map[string]interface{}{
		"name":      "Leg Day Program",
		"shared_by": "alice@example.com",
		"days": []map[string]interface{}{
			{"name": "Monday"},
			{"name": "Wednesday"},
		},
	}
	body, err := json.Marshal(payload)
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/api/v1/programs", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var created db.Program
	err = json.Unmarshal(w.Body.Bytes(), &created)
	assert.NoError(t, err)
	assert.NotEmpty(t, created.ID, "should have returned a new UUID")
	assert.Equal(t, "Leg Day Program", created.Name)
	assert.Equal(t, "alice@example.com", created.SharedBy)
	assert.Len(t, created.Days, 2)

	assert.WithinDuration(t, time.Now(), created.CreatedAt, time.Second*5)
	assert.WithinDuration(t, time.Now(), created.UpdatedAt, time.Second*5)

	w = httptest.NewRecorder()
	getURL := "/api/v1/programs/" + created.ID
	req, _ = http.NewRequest("GET", getURL, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var fetched db.Program
	err = json.Unmarshal(w.Body.Bytes(), &fetched)
	assert.NoError(t, err)

	assert.Equal(t, created.ID, fetched.ID)
	assert.Equal(t, created.Name, fetched.Name)
	assert.Equal(t, created.SharedBy, fetched.SharedBy)
	assert.Len(t, fetched.Days, 2)
	for _, d := range fetched.Days {
		assert.NotEmpty(t, d.ID)
		assert.Equal(t, fetched.ID, d.ProgramID)
		assert.NotEmpty(t, d.CreatedAt)
	}
}
