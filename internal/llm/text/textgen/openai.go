package textgen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/text/model"
)

type OpenAITextProvider struct {
	model  string
	url    string
	apiKey string
}

func NewOpenAITextProvider(apiKey string) (*OpenAITextProvider, error) {
	return &OpenAITextProvider{
		model:  "gpt-4o-mini",
		url:    "https://api.openai.com/v1/chat/completions",
		apiKey: apiKey,
	}, nil
}

func (p *OpenAITextProvider) GetModel() string {
	return p.model
}

func (p *OpenAITextProvider) GetURL() string {
	return p.url
}

func (p *OpenAITextProvider) GenerateImageDescription(
	role string,
	prompt string,
	conversationHistory []model.TextHistory,
) (string, error) {
	messages := []map[string]string{
		{
			"role":    "system",
			"content": role,
		},
	}

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
		"model":    p.model,
		"messages": messages,
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
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var llmResponse model.OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResponse); err != nil {
		return "", fmt.Errorf("error decoding response: %w", err)
	}

	if len(llmResponse.Choices) == 0 || llmResponse.Choices[0].Message.Content == "" {
		return "", fmt.Errorf("no content found in the response. %w\n", err)
	}

	return llmResponse.Choices[0].Message.Content, nil
}
