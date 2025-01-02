package textgen

import "fmt"

type ProviderType string

const (
	OpenAI       ProviderType = "openai"
	Anthropic    ProviderType = "anthropic"
	GoogleVertex ProviderType = "google-vertex"
)

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
