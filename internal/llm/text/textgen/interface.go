package textgen

import "github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/models"

type Provider interface {
	GetModel() string

	GetURL() string

	GenerateImageDescription(
		role string,
		prompt string,
		conversationHistory []models.TextHistory,
	) (string, error)
}
