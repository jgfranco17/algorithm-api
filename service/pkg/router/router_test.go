package router

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func parseJsonData(t *testing.T, r *httptest.ResponseRecorder) map[string]string {
	var responseBody map[string]string
	err := json.Unmarshal(r.Body.Bytes(), &responseBody)
	assert.NoError(t, err)
	return responseBody
}

func TestCreateValidServer(t *testing.T) {
	version := 1
	port := 8080

	server, serverErr := CreateServer(version, port)
	assert.NoError(t, serverErr)
	assert.NotNil(t, server)
	assert.NotNil(t, server.Router)
	assert.Equal(t, ":8080", server.Port)
}

func TestCreateServer(t *testing.T) {
	server, serverErr := CreateServer(-1, 8080)
	assert.Nil(t, server)
	assert.NotNil(t, serverErr)
}

func TestHomeHandler(t *testing.T) {
	server, serverErr := CreateServer(1, 8080)
	assert.NoError(t, serverErr)
	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	server.Router.ServeHTTP(resp, req)
	data := parseJsonData(t, resp)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "Welcome to the Algorithm API page!", data["message"])
}

func TestAboutHandler(t *testing.T) {
	server, serverErr := CreateServer(1, 8080)
	assert.NoError(t, serverErr)
	req, _ := http.NewRequest("GET", "/about", nil)
	resp := httptest.NewRecorder()
	server.Router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	data := parseJsonData(t, resp)
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "Chino Franco", data["author"])
	assert.Equal(t, "Algorithm API", data["title"])
}

func TestNotFoundHandler(t *testing.T) {
	server, serverErr := CreateServer(1, 8080)
	assert.NoError(t, serverErr)
	req, _ := http.NewRequest("GET", "/broken", nil)
	resp := httptest.NewRecorder()
	server.Router.ServeHTTP(resp, req)
	data := parseJsonData(t, resp)
	assert.Equal(t, 404, resp.Code)
	assert.Equal(t, "Endpoint not found", data["error"])
}

func TestFibonacciHandler(t *testing.T) {
	server, serverErr := CreateServer(1, 8080)
	assert.NoError(t, serverErr)
	req, _ := http.NewRequest("GET", "/v1/algorithms/fibonacci/5", nil)
	resp := httptest.NewRecorder()

	server.Router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestTwoSumHandler(t *testing.T) {
	server, serverErr := CreateServer(1, 8080)
	assert.NoError(t, serverErr)
	req, _ := http.NewRequest("GET", "/v1/algorithms/twosum/1-2-3/4", nil)
	resp := httptest.NewRecorder()

	server.Router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
