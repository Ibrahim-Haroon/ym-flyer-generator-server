package imagegen

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Ibrahim-Haroon/ym-flyer-generator-server.git/internal/llm/image/models"
	"image"
	"image/png"
	"net/http"
	"os"
	"path/filepath"
)

type OpenAIImageProvider struct {
	model  string
	url    string
	apiKey string
}

func NewOpenAIImageProvider() (*OpenAIImageProvider, error) {
	return &OpenAIImageProvider{
		model:  "dall-e-3",
		url:    "https://api.openai.com/v1/images/generations",
		apiKey: os.Getenv("OPENAI_API_KEY"),
	}, nil
}

func (p *OpenAIImageProvider) GetModel() string {
	return p.model
}

func (p *OpenAIImageProvider) GetURL() string {
	return p.url
}

func (p *OpenAIImageProvider) GenerateImage(imageDescription string) ([]string, error) {
	payload := map[string]any{
		"model":           p.model,
		"prompt":          imageDescription,
		"size":            "1024x1024",
		"response_format": "b64_json",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, p.url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var llmResponse models.OpenAIResponse
	if err := json.NewDecoder(resp.Body).Decode(&llmResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	if len(llmResponse.Data) == 0 {
		return nil, fmt.Errorf("no content found in the response")
	}

	return p.saveImages(llmResponse)
}

func (p *OpenAIImageProvider) saveImages(llmResponse models.OpenAIResponse) ([]string, error) {
	var savedPaths []string
	path := filepath.Join("images", fmt.Sprintf("%d", llmResponse.Created))
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("Failed to create directory: " + err.Error())
	}

	for i, data := range llmResponse.Data {
		imageBytes, err := base64.StdEncoding.DecodeString(data.B64_json)
		if err != nil {
			return nil, fmt.Errorf("failed to decode base64 string for image %d: %w", i, err)
		}

		filename := fmt.Sprintf("original-image-%d.png", i)
		savePath := filepath.Join(path, filename)

		img, _, err := image.Decode(bytes.NewReader(imageBytes))
		if err != nil {
			return nil, fmt.Errorf("failed to decode image %d: %w", i, err)
		}

		file, err := os.Create(savePath)
		if err != nil {
			return nil, fmt.Errorf("failed to create file %s: %w", savePath, err)
		}
		defer file.Close()

		if err := png.Encode(file, img); err != nil {
			return nil, fmt.Errorf("failed to encode image %d as PNG: %w", i, err)
		}

		savedPaths = append(savedPaths, savePath)
	}

	return savedPaths, nil
}
