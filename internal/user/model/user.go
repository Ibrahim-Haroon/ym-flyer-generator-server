package model

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

// @Description User information
type User struct {
	// Unique identifier for the user
	// Example: "123e4567-e89b-12d3-a456-426614174000"
	ID string `json:"id"`

	// Username for login
	// Example: "john_doe"
	Username string `json:"username"`

	// Hashed password
	PasswordHash string `json:"-"`

	// User's email address
	// Example: "john@example.com"
	Email string `json:"email"`

	// Account creation timestamp
	// Example: "2024-01-04T12:00:00Z"
	CreatedAt string `json:"created_at"`

	// Last update timestamp
	// Example: "2024-01-04T12:00:00Z"
	UpdatedAt string `json:"updated_at"`

	// Last login timestamp
	// Example: "2024-01-04T12:00:00Z"
	LastLogin string `json:"last_login"`

	// Whether the account is active
	// Example: true
	ActiveStatus bool `json:"active_status"`

	// Whether the user has admin privileges
	// Example: false
	IsAdmin bool `json:"is_admin"`

	// API keys for text generation services (encrypted)
	TextModelApiKeys map[textgen.ProviderType]string `json:"-"`

	// API keys for image generation services (encrypted)
	ImageModelApiKeys map[imagegen.ProviderType]string `json:"-"`
}

// @Description Response to getting availible LLM Providers
type AvailableLLMProvidersResponse struct {
	Providers map[string][]string `json:"providers"`
}

// @Description Request to add/update API keys
type UpdateLLMProviderAPIKeysRequest struct {
	// Text generation service provider
	// Example: "anthropic"
	TextProvider string `json:"text_provider" binding:"required"`

	// API key for text generation service
	// Example: "sk-..."
	TextAPIKey string `json:"text_api_key" binding:"required"`

	// Image generation service provider
	// Example: "openai"
	ImageProvider string `json:"image_provider" binding:"required"`

	// API key for image generation service
	// Example: "sk-..."
	ImageAPIKey string `json:"image_api_key" binding:"required"`
}

// @Description Response for updating user, such as deleting a user
type UpdateUserResponse struct {
	// Message regarding successful operation
	Message string `json:"message"`
}

// @Description Response for getting user, everything but password and api keys
type UserResponse struct {
	// Unique identifier for the user
	// Example: "123e4567-e89b-12d3-a456-426614174000"
	ID string `json:"id"`

	// Username for login
	// Example: "john_doe"
	Username string `json:"username"`

	// User's email address
	// Example: "john@example.com"
	Email string `json:"email"`

	// Account creation timestamp
	// Example: "2024-01-04T12:00:00Z"
	CreatedAt string `json:"created_at"`

	// Last update timestamp
	// Example: "2024-01-04T12:00:00Z"
	UpdatedAt string `json:"updated_at"`

	// Last login timestamp
	// Example: "2024-01-04T12:00:00Z"
	LastLogin string `json:"last_login"`

	// Whether the account is active
	// Example: true
	ActiveStatus bool `json:"active_status"`

	// Whether the user has admin privileges
	// Example: false
	IsAdmin bool `json:"is_admin"`
}
