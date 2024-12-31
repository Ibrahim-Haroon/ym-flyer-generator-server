package textgen

import (
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/models"
)

type GoogleVertexTextProvider struct {
	// TODO: Different struct than anthropci and openai
	// model, url, project id, location, response type
}

func NewGoogleVertexTextProvider() (*GoogleVertexTextProvider, error) {
	return nil, fmt.Errorf("Google Gemini support coming soon!")
}

func (p *GoogleVertexTextProvider) GetModel() string {
	return ""
}

func (p *GoogleVertexTextProvider) GetURL() string {
	return ""
}

func (p *GoogleVertexTextProvider) GenerateImageDescription(
	role string,
	prompt string,
	conversationHistory []models.TextHistory,
) (string, error) {
	return "", nil
}
