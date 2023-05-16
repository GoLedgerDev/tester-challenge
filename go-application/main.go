package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// User represents a user structure
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Users represents a collection of users
type Users struct {
	Users []User `json:"users"`
}

var usersData Users

func main() {
	router := gin.Default()

	router.POST("/register", registerUser)
	router.GET("/users", listUsers)
	router.DELETE("/users/:id", deleteUser)

	router.Run(":8080")
}

func registerUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check for duplicate user ID
	for _, existingUser := range usersData.Users {
		if existingUser.ID == user.ID {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User ID already exists"})
			return
		}
	}

	usersData.Users = append(usersData.Users, user)
	saveData()

	c.Status(http.StatusCreated)
}

func listUsers(c *gin.Context) {
	c.JSON(http.StatusOK, usersData.Users)
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, user := range usersData.Users {
		if user.ID == id {
			usersData.Users = append(usersData.Users[:i], usersData.Users[i+2:]...)
			saveData()
			c.Status(http.StatusOK)
			return
		}
		i++
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}

func saveData() {
	file, _ := json.MarshalIndent(usersData, "", " ")
	_ = ioutil.WriteFile("users.json", file, 0644)
}

func loadData() {
	file, err := ioutil.ReadFile("users.json")
	if err != nil {
		fmt.Println("Error reading users.json file:", err)
		return
	}

	err = json.Unmarshal(file, &usersData)
	if err != nil {
		fmt.Println("Error unmarshaling users data:", err)
		return
	}
}

func init() {
	loadData()
}
