package flyer

import (
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/flyer/model"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/template"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

func CreateBackground(
	colorPalette string,
	textModelProvider model.TextModelProviderMeta,
	imageModelProvider model.ImageModelProviderMeta,
) ([]string, error) {
	imageModel, err := imagegen.NewProvider(imageModelProvider.Name, imageModelProvider.ApiKey)
	if err != nil {
		return nil, fmt.Errorf("Error getting image model: " + err.Error())
	}

	imageDescriptionGenerationModel, err := textgen.NewProvider(textModelProvider.Name, imageModelProvider.ApiKey)
	if err != nil {
		return nil, fmt.Errorf("Error getting text model: " + err.Error())
	}

	imageDescription, err := imageDescriptionGenerationModel.GenerateImageDescription(
		template.Role,
		template.ImageDescriptonGenerationPrompt(colorPalette),
		nil, // no conversation history
	)
	if err != nil {
		return nil, fmt.Errorf("Error generating image description: " + err.Error())
	}

	backgroundSavePaths, err := imageModel.GenerateImage(imageDescription)
	if err != nil {
		return nil, fmt.Errorf("Error getting image generations: " + err.Error())
	}

	return backgroundSavePaths, nil
}
