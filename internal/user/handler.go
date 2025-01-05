package user

import (
	"net/http"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/encryption"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user/model"
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

// @Summary Register new user
// @Description Register a new user with username, password and email
// @Tags users
// @Accept json
// @Produce json
// @Param request body model.RegisterRequest true "User registration details"
// @Success 201 {object} model.RegisterResponse "Successfully registered"
// @Failure 400 {object} model.UserErrorResponse "Invalid request body or username taken"
// @Router /api/v1/users/register [post]
func (h *Handler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.UserErrorResponse{Error: "Invalid request body"})
		return
	}

	user := &model.User{
		Username:     req.Username,
		PasswordHash: req.Password,
		Email:        req.Email,
	}

	token, err := h.service.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.RegisterResponse{
		Token: token,
		User: model.UserRegistrationInfo{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept json
// @Produce json
// @Param request body model.LoginRequest true "Login credentials"
// @Success 200 {object} model.LoginResponse "Successfully logged in"
// @Failure 400 {object} model.UserErrorResponse "Invalid request body"
// @Failure 401 {object} model.UserErrorResponse "Invalid credentials"
// @Router /api/v1/users/login [post]
func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.UserErrorResponse{Error: "Invalid request body"})
		return
	}

	token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, model.UserErrorResponse{Error: "Invalid credentials"})
		return
	}

	user, err := h.service.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, model.UserErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{
		ID:    user.ID,
		Token: token,
	})
}

// @Summary Get user details by ID
// @Description Get detailed information about a user by ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} model.UserResponse "User details"
// @Failure 401 {object} model.UserErrorResponse "Unauthorized access"
// @Failure 403 {object} model.UserErrorResponse "Forbidden"
// @Failure 404 {object} model.UserErrorResponse "User not found"
// @Router /api/v1/users/{id} [get]
func (h *Handler) GetUserById(c *gin.Context) {
	userID := c.Param("id")

	userIDFromClaim, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.UserErrorResponse{Error: "No authentication claims found"})
		return
	}

	isAdmin, _ := c.Get("isAdmin")
	if userIDFromClaim.(string) != userID && !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, model.UserErrorResponse{Error: "Unauthorized access"})
		return
	}

	user, err := h.service.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, model.UserErrorResponse{Error: "User not found"})
		return
	}

	c.JSON(http.StatusOK, model.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		LastLogin:    user.LastLogin,
		ActiveStatus: user.ActiveStatus,
		IsAdmin:      user.IsAdmin,
	})
}

// @Summary Get the available LLM Providers for the user
// @Description Retrieve a list of available text and image llm providers
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} model.AvailableLLMProvidersResponse "Retrieved list of available LLM models"
// @Failure 400 {object} model.UserErrorResponse "Invalid request body"
// @Failure 401 {object} model.UserErrorResponse "Unauthorized access"
// @Router /api/v1/users/{id}/api-keys [get]
func (h *Handler) GetAvailableLLMProviders(c *gin.Context) {
	userID := c.Param("id")

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, model.UserErrorResponse{Error: "Unauthorized access"})
		return
	}

	availableLLMProviders, err := h.service.GetAvailableLLMProviders(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.UserErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.AvailableLLMProvidersResponse{
		Providers: availableLLMProviders,
	})
}

// @Summary Add/update LLM Providers API keys (idempotent)
// @Description Update user's API keys for text and image generation services
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param request body model.UpdateAPIKeysRequest true "API keys"
// @Security BearerAuth
// @Success 200 {object} model.UpdateUserResponse "API keys updated"
// @Failure 400 {object} model.UserErrorResponse "Invalid request body"
// @Failure 401 {object} model.UserErrorResponse "Unauthorized access"
// @Router /api/v1/users/{id}/api-keys [put]
func (h *Handler) UpdateLLMProviderAPIKeys(c *gin.Context) {
	userID := c.Param("id")

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, model.UserErrorResponse{Error: "Unauthorized access"})
		return
	}

	var req model.UpdateLLMProviderAPIKeysRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.UserErrorResponse{Error: "Invalid request body"})
		return
	}

	keys := &encryption.APIKeys{
		TextProviders:  req.TextProviders,
		ImageProviders: req.ImageProviders,
	}

	if err := h.service.UpdateLLMProviderAPIKeys(userID, keys); err != nil {
		c.JSON(http.StatusBadRequest, model.UserErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.UpdateUserResponse{
		Message: "API keys successfully updated",
	})
}

// @Summary Delete user
// @Description Delete a user account
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Security BearerAuth
// @Success 200 {object} model.UpdateUserResponse "User deleted"
// @Failure 401 {object} model.UserErrorResponse "Unauthorized access"
// @Failure 403 {object} model.UserErrorResponse "Forbidden"
// @Failure 404 {object} model.UserErrorResponse "User not found"
// @Router /api/v1/users/{id} [delete]
func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	userIDFromClaim, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, model.UserErrorResponse{Error: "No authentication claims found"})
		return
	}

	isAdmin, _ := c.Get("isAdmin")
	if userIDFromClaim.(string) != userID && !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, model.UserErrorResponse{Error: "Unauthorized access"})
		return
	}

	if err := h.service.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, model.UserErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.UpdateUserResponse{
		Message: "User deleted successfully",
	})
}
