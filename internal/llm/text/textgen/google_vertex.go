package textgen

import (
	"fmt"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/model"
)

type GoogleVertexTextProvider struct {
	// TODO: Different struct than anthropci and openai
	// model, url, project id, location, response type
}

func NewGoogleVertexTextProvider(apiKey string) (*GoogleVertexTextProvider, error) {
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
	conversationHistory []model.TextHistory,
) (string, error) {
	return "", nil
}
