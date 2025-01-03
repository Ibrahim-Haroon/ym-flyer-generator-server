package model

import (
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/imagegen"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/textgen"
)

type User struct {
	ID                string                           `json:"id"`
	Username          string                           `json:"username"`
	PasswordHash      string                           `json:"password_hash"`
	Email             string                           `json:"email"`
	IsAdmin           bool                             `json:"is_admin"`
	TextModelApiKeys  map[textgen.ProviderType]string  `json:"text_model_api_keys"`
	ImageModelApiKeys map[imagegen.ProviderType]string `json:"image_model_api_keys"`
}
