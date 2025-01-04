package jwt

import (
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user/model"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type Claims struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

type Service struct {
	signingKey []byte
}

func NewService(signingKey []byte) *Service {
	return &Service{
		signingKey: signingKey,
	}
}

func (s *Service) GenerateToken(user *model.User) (string, error) {
	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.signingKey)
}

func (s *Service) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.signingKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("error parsing token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}

	if !token.Valid {
		if time.Now().After(claims.ExpiresAt.Time) {
			return nil, fmt.Errorf("token has expired")
		}
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

func (s *Service) RefreshToken(tokenString string) (string, error) {
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return "", err
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.signingKey)
}
