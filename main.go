package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Role      string `json:"role"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	LastLogin string `json:"last_login"`
	Avatar    string `json:"avatar"`
}

var users = []user{{
	ID:        "1",
	Name:      "admin",
	Email:     "XXXXXXXXXXXXXXX",
	Role:      "admin",
	Status:    "active",
	CreatedAt: "XXXXXXXXXXXXXXX",
	UpdatedAt: "XXXXXXXXXXXXXXX",
	DeletedAt: "XXXXXXXXXXXXXXX",
	LastLogin: "XXXXXXXXXXXXXXX",
	Avatar:    "XXXXXXXXXXXXXXX",
}, {
	ID:        "2",
	Name:      "user",
	Email:     "XXXXXXXXXXXXXXX",
	Role:      "user",
	Status:    "active",
	CreatedAt: "XXXXXXXXXXXXXXX",
	UpdatedAt: "XXXXXXXXXXXXXXX",
	DeletedAt: "XXXXXXXXXXXXXXX",
	LastLogin: "XXXXXXXXXXXXXXX",
	Avatar:    "XXXXXXXXXXXXXXX",
}}

func getUsers(c *gin.Context) {
	
	c.IndentedJSON(http.StatusOK, users)
}

func getUserById(c *gin.Context) {
	id := c.Param("id")
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func createUser(c *gin.Context) {
	var newUser user

	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)

}

func deleteUSerById(c *gin.Context) {
	id := c.Param("id")

	for index, user := range users {
		if user.ID == id {
			users = append(users[:index], users[index+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "User deleted"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not found"})

}

func home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "App is running",
	})
}

func main() {
	router := gin.Default()
	router.GET("/", home)
	router.GET("/users", getUsers)
	router.POST("/users", createUser)
	router.GET("/users/:id", getUserById)
	router.DELETE("/users/:id", deleteUSerById)
	router.Run("localhost:8080")
}
