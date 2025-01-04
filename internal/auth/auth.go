package auth

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/auth/jwt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user/model"
)

type Service struct {
	jwtService *jwt.Service
}

func NewService(signingKey []byte) *Service {
	return &Service{
		jwtService: jwt.NewService(signingKey),
	}
}

func (s *Service) GenerateToken(user *model.User) (string, error) {
	return s.jwtService.GenerateToken(user)
}

func (s *Service) ValidateToken(tokenString string) (*jwt.Claims, error) {
	return s.jwtService.ValidateToken(tokenString)
}

func (s *Service) RefreshToken(tokenString string) (string, error) {
	return s.jwtService.RefreshToken(tokenString)
}
