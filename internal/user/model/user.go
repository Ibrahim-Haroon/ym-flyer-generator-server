package model

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

type User struct {
	ID                string                           `json:"id"`
	Username          string                           `json:"username"`
	PasswordHash      string                           `json:"password_hash"`
	Email             string                           `json:"email"`
	CreatedAt         string                           `json:"created_at"`
	UpdateAt          string                           `json:"updated_at"`
	LastLogin         string                           `json:"last_login"`
	ActiveStatus      bool                             `json:"active_status"`
	IsAdmin           bool                             `json:"is_admin"`
	TextModelApiKeys  map[textgen.ProviderType]string  `json:"text_model_api_keys"`
	ImageModelApiKeys map[imagegen.ProviderType]string `json:"image_model_api_keys"`
}

type UpdateAPIKeysRequest struct {
	TextProvider  string `json:"text_provider" binding:"required"`
	TextAPIKey    string `json:"text_api_key" binding:"required"`
	ImageProvider string `json:"image_provider" binding:"required"`
	ImageAPIKey   string `json:"image_api_key" binding:"required"`
}

type UpdateUserResponse struct {
	Message string `json:"message"`
}

type UserResponse struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
	LastLogin    string `json:"last_login"`
	ActiveStatus bool   `json:"active_status"`
	IsAdmin      bool   `json:"is_admin"`
}
