package textgen

import "fmt"

type ProviderType string

const (
	OpenAI       ProviderType = "openai"
	Anthropic    ProviderType = "anthropic"
	GoogleVertex ProviderType = "google-vertex"
)

func NewProvider(providerType ProviderType) (Provider, error) {
	switch providerType {
	case OpenAI:
		return NewOpenAITextProvider()
	case Anthropic:
		return NewAnthropicTextProvider()
	case GoogleVertex:
		return NewGoogleVertexTextProvider()
	default:
		return nil, fmt.Errorf("unsupported text provider: %s", providerType)
	}
}
