package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"core/internal/api"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// setupRouter initializes the Gin router with routes
func setupRouter() *gin.Engine {
	router := gin.Default()
	api.RegisterRoutes(router) // Register the application's routes
	return router
}

func TestRouteCore(t *testing.T) {
	// Initialize router
	router := setupRouter()

	// Mock request
	req, err := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Record the response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"service":"CORE","users":["user1","user2","user3"]}`, w.Body.String())
}
