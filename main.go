package main

import (
	"fmt"
	"log"
	"net/http"

	"golang-crud/initialzers"
	"github.com/gin-gonic/gin"
)

func init() {
	// Load environment variables
	initializers.LoadEnvVariables()

	// Connect to database
	initializers.ConnectDatabase()
}

func main() {
	r := gin.Default()

	// Health check
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Run server on port 8080
	fmt.Println("Server is running on port 8080...")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
