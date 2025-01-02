package model

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

type CreateRequest struct {
	// NumberOfFlyers uint `json:"number_of_flyers"`   TBD whether multiple file generations will be added
	ColorPalette           string                 `json:"color_palette"`
	TextModelProviderMeta  TextModelProviderMeta  `json:"text_model_provider"`
	ImageModelProviderMeta ImageModelProviderMeta `json:"image_model_provider"`
}

type TextModelProviderMeta struct {
	Name   textgen.ProviderType `json:"name"`
	ApiKey string               `json:"api_key"`
}

type ImageModelProviderMeta struct {
	Name   imagegen.ProviderType `json:"name"`
	ApiKey string                `json:"api_key"`
}

type CreateResponse struct {
	BackgroundPaths []string `json:"backgroud_paths"`
}
