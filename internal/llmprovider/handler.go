package llmprovider

import (
	"net/http"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llmprovider/model"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

// @Summary Get all the availible providers for a LLM type (text/image)
// @Description Retrieves all supported LLM providers for either image or text generation models
// @Tags provider
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param model_type either text or image
// @Security BearerAuth
// @Success 200 {file} binary "Image data"
// @Failure 401 {object} model.LlmProviderErrorResponse "Unauthorized access"
// @Failure 400 {object} model.LlmProviderErrorResponse "If the model type is not image or text"
// @Failure 500 {object} model.LlmProviderErrorResponse "Server error"
// @Router /api/v1/provider/{id}/{model_type} [get]
func (h *Handler) ListLLMProviders(c *gin.Context) {
	userID := c.Param("id")
	llmType := c.Param("llm_type")

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, model.LLMProviderErrorResponse{Error: "Unauthorized access"})
		return
	}

	providers, err := h.service.GetLLMProviders(llmType)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.LLMProviderErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.LLMProviders{Providers: providers})
}
