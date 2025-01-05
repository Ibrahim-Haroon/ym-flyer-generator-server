package llmprovider

import (
	"fmt"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) GetLLMProviders(llmType string) ([]string, error) {
	if llmType == "text" {
		return textgen.GetAllProviders(), nil
	}

	if llmType == "image" {
		return imagegen.GetAllProviders(), nil
	}

	return nil, fmt.Errorf("Expected LLM type to be either \"text\" \"image\", got %s", llmType)
}
