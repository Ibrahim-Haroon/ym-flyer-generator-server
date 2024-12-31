package provider

import "fmt"

type TextProviderType string

const (
	OpenAI       TextProviderType = "openai"
	Anthropic    TextProviderType = "anthropic"
	GoogleVertex TextProviderType = "google-vertex"
)

func NewProvider(textProviderType TextProviderType) (TextProvider, error) {
	switch textProviderType {
	case OpenAI:
		return NewOpenAITextProvider()
	case Anthropic:
		return NewAnthropicTextProvider()
	case GoogleVertex:
		return NewGoogleVertexTextProvider()
	default:
		return nil, fmt.Errorf("unsupported text provider: %s", textProviderType)
	}
}
