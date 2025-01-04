package imagegen

import (
	"errors"
	"fmt"
)

type ProviderType string

const (
	OpenAI ProviderType = "openai"
)

var validProviderTypes = map[ProviderType]bool{
	OpenAI: true,
}

func NewProviderType(value string) (ProviderType, error) {
	pt := ProviderType(value)
	if !validProviderTypes[pt] {
		return "", errors.New(fmt.Sprintf("%s is not a supported image model", value))
	}
	return pt, nil
}

func NewProvider(providerType ProviderType, apiKey string) (Provider, error) {
	switch providerType {
	case OpenAI:
		return NewOpenAIImageProvider(apiKey)
	default:
		return nil, fmt.Errorf("unsupported image provider: %s", providerType)
	}
}
