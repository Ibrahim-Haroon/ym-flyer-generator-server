package model

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

type CreateRequest struct {
	// NumberOfFlyers uint `json:"number_of_flyers"`   TBD whether multiple file generations will be added
	ColorPalette       string                `json:"color_palette"`
	TextModelProvider  textgen.ProviderType  `json:"text_model_provider"  binding:"required"`
	ImageModelProvider imagegen.ProviderType `json:"image_model_provider" binding:"required"`
}

type CreateResponse struct {
	BackgroundPaths []string `json:"background_paths"`
}
