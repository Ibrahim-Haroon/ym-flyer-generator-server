package main

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/flyer"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func HealthCheck(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "healthy")
}

func main() {
	err := godotenv.Load("local-config.env")
	if err != nil {
		log.Fatal("Error loading in enviornment file!")
	}

	router := gin.Default()
	router.GET("/api/v1/health", HealthCheck)
	router.GET("/api/v1/flyer/*path", flyer.GetBackground)
	router.POST("/api/v1/flyer", flyer.GenerateBackgrounds)

	if os.Getenv("CONTAINER") == "TRUE" {
		router.Run(":8080")
	} else {
		router.Run("localhost:8080")
	}
}
