package flyer

import (
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/template"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/user"
)

type Service struct {
	userService *user.Service
}

func NewService(userService *user.Service) *Service {
	return &Service{
		userService: userService,
	}
}

func (s *Service) CreateBackground(
	userID string,
	colorPalette string,
	textModelProvider textgen.ProviderType,
	imageModelProvider imagegen.ProviderType,
) ([]string, error) {
	apiKeys, err := s.userService.GetDecryptedAPIKeys(userID, textModelProvider, imageModelProvider)
	if err != nil {
		return nil, fmt.Errorf("failed to get API keys: %w", err)
	}

	imageModel, err := imagegen.NewProvider(apiKeys.ImageProvider, apiKeys.ImageAPIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize image provider: %w", err)
	}

	textModel, err := textgen.NewProvider(apiKeys.TextProvider, apiKeys.TextAPIKey)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize text provider: %w", err)
	}

	imageDescription, err := textModel.GenerateImageDescription(
		template.Role,
		template.ImageDescriptonGenerationPrompt(colorPalette),
		nil, // no conversation history
	)
	if err != nil {
		return nil, fmt.Errorf("failed to generate image description: %w", err)
	}

	backgroundSavePaths, err := imageModel.GenerateImage(imageDescription)
	if err != nil {
		return nil, fmt.Errorf("failed to generate image: %w", err)
	}

	return backgroundSavePaths, nil
}
