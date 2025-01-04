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

func (h *Handler) Register(c *gin.Context) {
	var req model.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user := &model.User{
		Username:     req.Username,
		PasswordHash: req.Password,
		Email:        req.Email,
	}

	token, err := h.service.Register(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.RegisterResponse{
		Token: token,
		User: model.UserRegisterationInfo{
			UserID:   user.ID,
			Username: user.Username,
			Email:    user.Email,
		},
	})
}

func (h *Handler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := h.service.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	user, err := h.service.GetUserByUsername(req.Username)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, model.LoginResponse{
		ID:    user.ID,
		Token: token,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	userID := c.Param("id")

	userIDFromClaim, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication claims found"})
		return
	}

	isAdmin, _ := c.Get("isAdmin")
	if userIDFromClaim.(string) != userID && !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		return
	}

	user, err := h.service.GetUserById(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, model.UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Email:        user.Email,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdateAt,
		LastLogin:    user.LastLogin,
		ActiveStatus: user.ActiveStatus,
		IsAdmin:      user.IsAdmin,
	})
}

func (h *Handler) UpdateAPIKeys(c *gin.Context) {
	userID := c.Param("id")

	userIDFromClaim, exists := c.Get("userID")
	if !exists || userIDFromClaim.(string) != userID {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
		return
	}

	var req model.UpdateAPIKeysRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	keys := &encryption.APIKeys{
		TextProvider:  req.TextProvider,
		TextAPIKey:    req.TextAPIKey,
		ImageProvider: req.ImageProvider,
		ImageAPIKey:   req.ImageAPIKey,
	}

	if err := h.service.UpdateAPIKeys(userID, keys); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.UpdateUserResponse{
		Message: "API keys successfully updated",
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	userID := c.Param("id")

	userIDFromClaim, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "No authentication claims found"})
		return
	}

	isAdmin, _ := c.Get("isAdmin")
	if userIDFromClaim.(string) != userID && !isAdmin.(bool) {
		c.JSON(http.StatusForbidden, gin.H{"error": "Unauthorized access"})
		return
	}

	if err := h.service.DeleteUser(userID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.UpdateUserResponse{
		Message: "User deleted successfully",
	})
}
