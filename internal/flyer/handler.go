package flyer

import (
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/flyer/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func GetBackground(c *gin.Context) {
	imagePath := c.Param("path")
	fullPath := filepath.Join(".", imagePath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	imageData, err := os.ReadFile(fullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to read image: %v", err)})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", imageData)
}

func GenerateBackgrounds(c *gin.Context) {
	var createRequest model.CreateRequest

	err := c.BindJSON(&createRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("failed to bind json: %v", err))
	}

	backgroundPaths, err := CreateBackground(
		createRequest.ColorPalette,
		createRequest.TextModelProviderMeta,
		createRequest.ImageModelProviderMeta,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}

	c.JSON(http.StatusOK, model.CreateResponse{BackgroundPaths: backgroundPaths})
}
