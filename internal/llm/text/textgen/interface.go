package textgen

import "github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/model"

type Provider interface {
	GetModel() string

	GetURL() string

	GenerateImageDescription(
		role string,
		prompt string,
		conversationHistory []model.TextHistory,
	) (string, error)
}
