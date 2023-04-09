// api.go

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{1, "John Doe", "john@example.com"},
	{2, "Jane Doe", "jane@example.com"},
	{3, "Bob Smith", "bob@example.com"},
}

func init() {
	router := gin.Default()
	router.GET("/users", GetUsers)
	router.Run(":8080")
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}
