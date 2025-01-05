package textgen

import (
	"errors"
	"fmt"
)

type ProviderType string

const (
	OpenAI       ProviderType = "openai"
	Anthropic    ProviderType = "anthropic"
	GoogleVertex ProviderType = "google-vertex"
)

var validProviderTypes = map[ProviderType]bool{
	OpenAI:       true,
	Anthropic:    true,
	GoogleVertex: true,
}

func GetAllProviders() []string {
	providers := make([]string, 0, len(validProviderTypes))
	for provider := range validProviderTypes {
		providers = append(providers, string(provider))
	}
	return providers
}

func NewProviderType(value string) (ProviderType, error) {
	pt := ProviderType(value)
	if !validProviderTypes[pt] {
		return "", errors.New(fmt.Sprintf("%s is not a supported text model", value))
	}
	if pt == "google-vertex" {
		return "", errors.New("support for google gemini coming soon...")
	}
	return pt, nil
}

func NewProvider(providerType ProviderType, apiKey string) (Provider, error) {
	switch providerType {
	case OpenAI:
		return NewOpenAITextProvider(apiKey)
	case Anthropic:
		return NewAnthropicTextProvider(apiKey)
	case GoogleVertex:
		return NewGoogleVertexTextProvider(apiKey)
	default:
		return nil, fmt.Errorf("unsupported text provider: %s\n", providerType)
	}
}
