package provider

import "github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/models"

type TextProvider interface {
	GetModel() string

	GetURL() string

	GenerateImageDescription(
		role string,
		prompt string,
		conversationHistory []models.TextHistory,
	) (string, error)
}
