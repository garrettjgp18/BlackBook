package ai

import (
	"BlackBook/internal/terminal"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// OllamaClient handles communication with local Ollama
type OllamaClient struct {
	baseURL string
	model   string
	client  *http.Client
}

// OllamaRequest represents a request to Ollama
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

// OllamaResponse represents a response from Ollama
type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

// NewOllamaClient creates a new Ollama client
func NewOllamaClient(baseURL, model string) *OllamaClient {
	if baseURL == "" {
		baseURL = "http://localhost:11434"
	}
	if model == "" {
		model = "llama3"
	}

	return &OllamaClient{
		baseURL: baseURL,
		model:   model,
		client: &http.Client{
			Timeout: 120 * time.Second,
		},
	}
}

// IsAvailable checks if Ollama is running and accessible
func (c *OllamaClient) IsAvailable() bool {
	resp, err := c.client.Get(c.baseURL + "/api/tags")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == http.StatusOK
}

// Generate generates a response from Ollama
func (c *OllamaClient) Generate(prompt string) (string, error) {
	reqBody := OllamaRequest{
		Model:  c.model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	resp, err := c.client.Post(
		c.baseURL+"/api/generate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("ollama returned status %d: %s", resp.StatusCode, string(body))
	}

	var ollamaResp OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&ollamaResp); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	return ollamaResp.Response, nil
}

// AnalyzeCommand analyzes a specific command in context
func (c *OllamaClient) AnalyzeCommand(cmd terminal.CommandEntry, context []terminal.CommandEntry) (string, error) {
	prompt := BuildCommandAnalysisPrompt(cmd, context)
	return c.Generate(prompt)
}

// SuggestNextStep suggests the next pentest action
func (c *OllamaClient) SuggestNextStep(context []terminal.CommandEntry) (string, error) {
	prompt := BuildNextStepPrompt(context)
	return c.Generate(prompt)
}

// GenerateScript generates a custom script based on context
func (c *OllamaClient) GenerateScript(request string, context []terminal.CommandEntry) (string, error) {
	prompt := BuildScriptGenerationPrompt(request, context)
	return c.Generate(prompt)
}

// ExplainOutput explains command output
func (c *OllamaClient) ExplainOutput(output string) (string, error) {
	prompt := BuildOutputExplanationPrompt(output)
	return c.Generate(prompt)
}
