package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHomeHandler(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Welcome to the Algorithm API page!")
}

func TestAboutHandler(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/about")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Chino Franco")
	assert.Contains(t, w.Body.String(), "Algorithm API")
	assert.Contains(t, w.Body.String(), "https://github.com/jgfranco17/algorithm-api")
}

func TestNotFoundHandler(t *testing.T) {
	router := setupRouter()
	w := performRequest(router, "GET", "/nonexistent")
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "Endpoint not found")
}

// Helper functions for testing

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", HomeHandler)
	router.GET("/about", AboutHandler)
	router.NoRoute(NotFoundHandler)
	return router
}

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
