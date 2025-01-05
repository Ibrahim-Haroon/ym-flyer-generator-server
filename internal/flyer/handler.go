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

// @Summary Get a generated flyer background
// @Description Retrieves a previously generated flyer background image by file path
// @Tags flyer
// @Accept json
// @Produce image/png
// @Param id path string true "User ID"
// @Param path path string true "Image path"
// @Security BearerAuth
// @Success 200 {file} binary "Image data"
// @Failure 401 {object} model.GenerationErrorResponse "Unauthorized access"
// @Failure 404 {object} model.GenerationErrorResponse "Image not found"
// @Failure 500 {object} model.GenerationErrorResponse "Server error"
// @Router /api/v1/flyer/{id}/{path} [get]
func (h *Handler) GetBackground(c *gin.Context) {
	userID := c.Param("id")
	imagePath := c.Param("path")

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, model.GenerationErrorResponse{Error: "Unauthorized access"})
		return
	}

	fullPath := filepath.Join(".", imagePath)

	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, model.GenerationErrorResponse{Error: "Image not found"})
		return
	}

	imageData, err := os.ReadFile(fullPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GenerationErrorResponse{Error: fmt.Sprintf("Failed to read image: %v", err)})
		return
	}

	c.Header("Content-Type", "image/png")
	c.Data(http.StatusOK, "image/png", imageData)
}

// @Summary Generate new flyer backgrounds
// @Description Generates new flyer background images using AI with specified color palette
// @Tags flyer
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body model.CreateRequest true "Generation parameters"
// @Security BearerAuth
// @Success 202 {object} model.CreateResponse "Successfully queued generation"
// @Failure 400 {object} model.GenerationErrorResponse "Invalid request parameters"
// @Failure 401 {object} model.GenerationErrorResponse "Unauthorized access"
// @Failure 500 {object} model.GenerationErrorResponse "Server error"
// @Router /api/v1/flyer/{id} [post]
func (h *Handler) CreateBackground(c *gin.Context) {
	userID := c.Param("id")
	var createRequest model.CreateRequest

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, model.GenerationErrorResponse{Error: "Unauthorized access"})
		return
	}
	userID, ok := userIDFromClaim.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, model.GenerationErrorResponse{Error: "Type assertion on userIDFromClaim failed!"})
	}

	if err := c.BindJSON(&createRequest); err != nil {
		c.JSON(http.StatusBadRequest, model.GenerationErrorResponse{Error: fmt.Sprintf("Failed to bind request: %v", err)})
		return
	}

	backgroundPaths, err := h.service.CreateBackground(
		userID,
		createRequest.ColorPalette,
		createRequest.TextModelProvider,
		createRequest.ImageModelProvider,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.GenerationErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, model.CreateResponse{BackgroundPaths: backgroundPaths})
}
