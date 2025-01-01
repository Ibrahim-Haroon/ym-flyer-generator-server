package model

// DTO for openai image endpoint
type OpenAIResponse struct {
	Created int64  `json:"created"`
	Data    []Data `json:"data"`
}

type Data struct {
	B64JSON string `json:"b64_json"`
}
