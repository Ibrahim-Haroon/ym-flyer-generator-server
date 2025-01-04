package model

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=50"`
	Password string `json:"password" binding:"required,min=8"`
	Email    string `json:"email" binding:"required,email"`
}

type RegisterResponse struct {
	Token string                `json:"token"`
	User  UserRegisterationInfo `json:"user"`
}

type UserRegisterationInfo struct {
	UserID   string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
