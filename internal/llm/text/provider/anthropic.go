package provider

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/models"
	"net/http"
	"os"
)

type AnthropicTextProvider struct {
	model  string
	url    string
	apiKey string
}

func NewAnthropicTextProvider() (*AnthropicTextProvider, error) {
	return &AnthropicTextProvider{
		model:  "claude-3-sonnet-20240229",
		url:    "https://api.anthropic.com/v1/messages",
		apiKey: os.Getenv("ANTHROPIC_API_KEY"),
	}, nil
}

func (p *AnthropicTextProvider) GetModel() string {
	return p.model
}

func (p *AnthropicTextProvider) GetURL() string {
	return p.url
}

func (p *AnthropicTextProvider) GenerateImageDescription(
	role string,
	prompt string,
	conversationHistory []models.TextHistory,
) (string, error) {
	var messages []map[string]string

	if conversationHistory != nil {
		for _, msg := range conversationHistory {
			messages = append(messages, map[string]string{
				"role":    string(msg.Role),
				"content": msg.Content,
			})
		}
	}

	messages = append(messages, map[string]string{
		"role":    "user",
		"content": prompt,
	})

	payload := map[string]any{
		"model":      p.model,
		"system":     role,
		"messages":   messages,
		"max_tokens": 1_000,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, p.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("anthropic-version", "2023-06-01")
	req.Header.Set("x-api-key", p.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var llmResponse models.AnthropicResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResponse); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	if len(llmResponse.Content) == 0 {
		return "", fmt.Errorf("no content found in the response")
	}

	return llmResponse.Content[0].Text, nil
}
