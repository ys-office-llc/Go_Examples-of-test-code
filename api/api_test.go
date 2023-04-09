// api_test.go

package api_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/ys-office-llc/Go_Examples-of-test-code/api"
)

func TestGetUsers(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)

	router := setupRouter()
	router.ServeHTTP(w, req)

	var got []api.User
	err := json.Unmarshal(w.Body.Bytes(), &got)
	if err != nil {
		t.Errorf("Failed to unmarshal response body")
	}

	want := []api.User{
		{1, "John Doe", "john@example.com"},
		{2, "Jane Doe", "jane@example.com"},
		{3, "Bob Smith", "bob@example.com"},
	}

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, want, got)
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", api.GetUsers)
	return router
}
