package flyer

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/flyer/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) GetBackground(c *gin.Context) {
	userID := c.Param("id")
	imagePath := c.Param("path")

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

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

func (h *Handler) GenerateBackgrounds(c *gin.Context) {
	userID := c.Param("id")
	var createRequest model.CreateRequest

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}
	userID, ok := userIDFromClaim.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Type assertion on userIDFromClaim failed!"})
	}

	if err := c.BindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Failed to bind request: %v", err)})
		return
	}

	backgroundPaths, err := h.service.CreateBackground(
		userID,
		createRequest.ColorPalette,
		createRequest.TextModelProvider,
		createRequest.ImageModelProvider,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, model.CreateResponse{BackgroundPaths: backgroundPaths})
}
