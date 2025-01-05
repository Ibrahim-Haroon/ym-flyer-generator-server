package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/docs/swagger"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/flyer"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llmprovider"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/service"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func setupRouter(serviceModule *module.Module) *gin.Engine {
	router := gin.Default()

	userHandler := user.NewHandler(serviceModule.UserService)
	flyerHandler := flyer.NewHandler(serviceModule.FlyerService)
	llmProviderHandler := llmprovider.NewHandler(serviceModule.LlmProviderService)

	api := router.Group("/api/v1")
	{
		// Public endpoints (no auth required)
		api.GET("/health", HealthCheck)
		api.POST("/users/register", userHandler.Register)
		api.POST("/users/login", userHandler.Login)

		users := api.Group("/users")
		users.Use(serviceModule.Middleware.AuthUser)
		{
			users.GET("/:id", userHandler.GetUserById)
			users.GET("/:id/api-keys", userHandler.GetAvailableLLMProviders)
			users.PUT("/:id/api-keys", userHandler.UpdateLLMProviderAPIKeys)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		flyers := api.Group("/flyer")
		flyers.Use(serviceModule.Middleware.AuthUser)
		{
			flyers.POST("/:id", flyerHandler.CreateBackground)
			flyers.GET("/:id/*path", flyerHandler.GetBackground)
		}

		llmproviders := api.Group("/llm_provider")
		llmproviders.Use(serviceModule.Middleware.AuthUser)
		{
			llmproviders.GET("/:id/:llm_type", llmProviderHandler.GetLLMProviders)
		}

		admin := api.Group("/admin")
		admin.Use(serviceModule.Middleware.AuthAdmin)
		{
			// admin routes
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func main() {
	if err := godotenv.Load("local-config.env"); err != nil {
		log.Println("Error loading environment file! OK if running on AWS ECS")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	ctx := context.Background()
	serviceModule, err := module.NewModule(ctx, db)
	if err != nil {
		log.Fatal("Failed to initialize service module:", err)
	}
	defer serviceModule.Cleanup(ctx)

	router := setupRouter(serviceModule)

	port := ":8080"
	if os.Getenv("CONTAINER") != "TRUE" {
		log.Println("Application running inside container...")
		port = "localhost:8080"
	} else {
		log.Println("Application running locally...")
	}

	if err := router.Run(port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
