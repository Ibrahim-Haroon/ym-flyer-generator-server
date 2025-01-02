package model

// DTO for anthropic text endpoint
type AnthropicResponse struct {
	Content []Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
	Type string `json:"type"`
}
