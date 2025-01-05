package model

// @Description Registration request payload
type RegisterRequest struct {
	// Desired username
	// Example: "john_doe"
	Username string `json:"username" binding:"required,min=3,max=50"`

	// Password hash
	// Example: "$2a$10$..."
	Password string `json:"password" binding:"required,min=8"`

	// Email address
	// Example: "john@example.com"
	Email string `json:"email" binding:"required,email"`
}

// @Description Registration request payload
type RegisterResponse struct {
	// Token for all endpoints required auth
	Token string `json:"token"`
	// Basic Info regarding User
	User UserRegistrationInfo `json:"user"`
}

// @Description Basic Info regarding User
type UserRegistrationInfo struct {
	// user's UUID
	// Example: "john_doe"
	UserID string `json:"id"`

	// Desired username
	// Example: "john_doe"
	Username string `json:"username"`

	// Email address
	// Example: "john@example.com"
	Email string `json:"email"`
}
