package module

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/auth"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/encryption"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/flyer"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llmprovider"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/middleware"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user"
	"github.com/gin-gonic/gin"
)

type Module struct {
	db                 *sql.DB
	UserService        *user.Service
	FlyerService       *flyer.Service
	AuthService        *auth.Service
	LlmProviderService *llmprovider.Service
	Middleware         *MiddlewareModule
}

type MiddlewareModule struct {
	AuthUser  gin.HandlerFunc
	AuthAdmin gin.HandlerFunc
}

func NewModule(ctx context.Context, db *sql.DB) (*Module, error) {
	encryptionKey := []byte(os.Getenv("ENCRYPTION_KEY"))
	if len(encryptionKey) != 32 {
		return nil, fmt.Errorf("ENCRYPTION_KEY must be exactly 32 bytes")
	}

	encryptionService, err := encryption.NewService(encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize encryption service: %w", err)
	}

	signingKey := []byte(os.Getenv("JWT_SIGNING_KEY"))
	if len(signingKey) == 0 {
		return nil, fmt.Errorf("JWT_SIGNING_KEY environment variable is required")
	}

	userRepo := user.NewRepository(db)

	authService := auth.NewService(signingKey)
	userService := user.NewService(userRepo, authService, encryptionService)
	flyerService := flyer.NewService(userService)
	llmProviderService := llmprovider.NewService()

	middlewareModule := &MiddlewareModule{
		AuthUser:  middleware.AuthUserMiddleware(authService),
		AuthAdmin: middleware.AdminAuthMiddleware(authService),
	}

	return &Module{
		db:                 db,
		UserService:        userService,
		FlyerService:       flyerService,
		AuthService:        authService,
		LlmProviderService: llmProviderService,
		Middleware:         middlewareModule,
	}, nil
}

func (m *Module) Cleanup(ctx context.Context) error {
	if err := m.db.Close(); err != nil {
		return fmt.Errorf("error closing database connection: %w", err)
	}
	return nil
}
