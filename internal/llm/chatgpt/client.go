package chatgpt

import (
	"context"
	"fmt"
	"strings"

	"github.com/sashabaranov/go-openai"
)

// Client is the ChatGPT API client that implements the Paraphraser interface.
type Client struct {
	openAiClient *openai.Client
	cfg          *Config
}

func NewClient(cfg *Config) *Client {
	return &Client{
		openAiClient: openai.NewClient(cfg.ApiKey),
		cfg:          cfg,
	}
}

// Paraphrase builds a prompt to paraphrase the given text to sound more like the given tone
// and sends the corresponding request to ChatGPT API.
func (c *Client) Paraphrase(ctx context.Context, tone string, text string) (string, error) {
	resp, err := c.openAiClient.CreateChatCompletion(
		ctx,
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: buildParaphrasePrompt(tone, text),
				},
			},
			Temperature: c.cfg.Temperature,
		},
	)
	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}

func buildParaphrasePrompt(tone string, text string) string {
	return strings.TrimSpace(fmt.Sprintf(`
	Paraphrase the following text delimited by @@ to sound more %s.
	Do not include the @@ delimiters in the response.
	@@%s@@
	`, tone, text))
}
