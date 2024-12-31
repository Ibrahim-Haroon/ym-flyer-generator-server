package imagegen

import "fmt"

type ProviderType string

const (
	OpenAI ProviderType = "openai"
)

func NewProvider(providerType ProviderType) (Provider, error) {
	switch providerType {
	case OpenAI:
		return NewOpenAIImageProvider()
	default:
		return nil, fmt.Errorf("unsupported image provider: %s", providerType)
	}
}
