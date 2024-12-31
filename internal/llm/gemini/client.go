package gemini

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

// Client is the Gemini API client that implements the Paraphraser interface.
type Client struct {
	genModel *genai.GenerativeModel
}

func NewClient(cfg *Config) *Client {
	genAiClient, err := genai.NewClient(context.Background(), option.WithAPIKey(cfg.ApiKey))
	if err != nil {
		log.Fatal(err)
	}
	genModel := genAiClient.GenerativeModel("gemini-1.5-flash")
	genModel.SetTemperature(cfg.Temperature)

	return &Client{genModel: genModel}
}

// Paraphrase builds a prompt to paraphrase the given text to sound more like the given tone
// and sends the corresponding request to Gemini API.
func (c *Client) Paraphrase(ctx context.Context, tone string, text string) (string, error) {
	resp, err := c.genModel.GenerateContent(ctx, genai.Text(buildParaphrasePrompt(tone, text)))
	if err != nil {
		return "", fmt.Errorf("gemini generate content: %w", err)
	}

	if resp.PromptFeedback != nil && resp.PromptFeedback.BlockReason > 0 {
		return "", fmt.Errorf("prompt was blocked: %s", resp.PromptFeedback.BlockReason)
	}

	result := fmt.Sprintf("%s", resp.Candidates[0].Content.Parts[0])

	return result, nil
}

func buildParaphrasePrompt(tone string, text string) string {
	return strings.TrimSpace(fmt.Sprintf(`
	Paraphrase the following text to sound more %s.
	Only include the actual paraphrased text without surrounding quotes in the response.
	Try to keep the structure similar to that of the original text.
	Text: %s
	`, tone, text))
}
