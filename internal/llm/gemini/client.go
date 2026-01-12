package gemini

import (
	"context"
	"fmt"
	"log"
	"strings"

	"google.golang.org/genai"
)

// Client is the Gemini API client that implements the Paraphraser interface.
type Client struct {
	genaiClient   *genai.Client
	genContentCfg *genai.GenerateContentConfig
}

func NewClient(cfg *Config) *Client {
	genaiClient, err := genai.NewClient(context.Background(), &genai.ClientConfig{
		APIKey:  cfg.ApiKey,
		Backend: genai.BackendGeminiAPI,
	})
	if err != nil {
		log.Fatal(err)
	}

	genContentCfg := &genai.GenerateContentConfig{
		SystemInstruction: genai.NewContentFromText(systemPrompt, ""),
		Temperature:       genai.Ptr(cfg.Temperature),
	}

	return &Client{
		genaiClient:   genaiClient,
		genContentCfg: genContentCfg,
	}
}

// Paraphrase builds a prompt to paraphrase the given text to sound more like the given tone
// and sends the corresponding request to Gemini API.
func (c *Client) Paraphrase(ctx context.Context, tone string, text string) (string, error) {
	resp, err := c.genaiClient.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text(buildParaphrasePrompt(tone, text)),
		c.genContentCfg,
	)
	if err != nil {
		return "", fmt.Errorf("gemini generate content: %w", err)
	}

	if resp.PromptFeedback != nil && resp.PromptFeedback.BlockReason != "" {
		return "", fmt.Errorf("prompt was blocked: %s", resp.PromptFeedback.BlockReason)
	}

	return resp.Text(), nil
}

const systemPrompt = `
You are an expert paraphraser.
Your task is to rephrase the given text to match the specified tone while preserving its original meaning.
Only respond with the paraphrased output.
`

func buildParaphrasePrompt(tone string, text string) string {
	return strings.TrimSpace(fmt.Sprintf(`
Tone: %s
Text: %s
	`, tone, text))
}
