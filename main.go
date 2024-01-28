package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type user struct {
	ID    string
	Name  string
	Email string
}

var users = []user{
	{ID: "1", Name: "Fendy", Email: "fendy.baba@icloud.com"},
	{ID: "2", Name: "Lucas", Email: "lucas.vavon@icloud.com"},
	{ID: "3", Name: "Melvin", Email: "melvin.robin@icloud.com"},
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", postUsers)

	router.Run("localhost:8080")
}

// getUsers responds with the list of all users as JSON.
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
}

func postUsers(c *gin.Context) {
	var newUser user

	// Call BindJSON to bind the received JSON to
	// newUser.
	if err := c.BindJSON(&newUser); err != nil {
		return
	}

	// Add the new user to the slice.
	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// getUserByID locates the user whose ID value matches the id
// parameter sent by the client, then returns that user as a response.
func getUserByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of users, looking for
	// an user whose ID value matches the parameter.
	for _, a := range users {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
}
