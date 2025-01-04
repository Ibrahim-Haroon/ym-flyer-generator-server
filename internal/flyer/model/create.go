package model

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

// @Description Parameters for flyer background generation
type CreateRequest struct {
	// Color palette to use for the background
	// Example: "metalic gray and emerald green"
	ColorPalette string `json:"color_palette"`

	// Text generation model provider (e.g. "openai", "anthropic", "google-vertex")
	// Required: true
	// Example: "openai"
	TextModelProvider textgen.ProviderType `json:"text_model_provider" binding:"required"`

	// Image generation model provider (e.g. "openai")
	// Required: true
	// Example: "openai"
	ImageModelProvider imagegen.ProviderType `json:"image_model_provider" binding:"required"`
}

// @Description Response containing paths to generated backgrounds
type CreateResponse struct {
	// Array of file paths to the generated background images
	// Example: ["/images/1234567890/image.png"]
	BackgroundPaths []string `json:"background_paths"`
}
