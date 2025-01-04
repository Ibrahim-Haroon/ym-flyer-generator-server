package model

// @Description Login request payload
type LoginRequest struct {
	// Username for authentication
	// Example: "john_doe"
	Username string `json:"username" binding:"required"`

	// Password hash
	// Example: "$2a$10$..."
	Password string `json:"password" binding:"required"`
}

// @Description Login response
type LoginResponse struct {
	// User's unique identifier
	// Example: "123e4567-e89b-12d3-a456-426614174000"
	ID string `json:"id"`

	// JWT token for authentication
	// Example: "eyJhbGciOiJIUzI1NiIs..."
	Token string `json:"token"`
}
