package chatgpt

import (
	"context"
	"fmt"
	"strings"

	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
	"github.com/openai/openai-go/v3/responses"
)

// Client is the ChatGPT API client that implements the Paraphraser interface.
type Client struct {
	openaiClient openai.Client
}

func NewClient(cfg *Config) *Client {
	return &Client{
		openaiClient: openai.NewClient(
			option.WithAPIKey(cfg.ApiKey),
		),
	}
}

// Paraphrase builds a prompt to paraphrase the given text to sound more like the given tone
// and sends the corresponding request to ChatGPT API.
func (c *Client) Paraphrase(ctx context.Context, tone string, text string) (string, error) {
	params := responses.ResponseNewParams{
		Model:      openai.ChatModelGPT5Nano,
		Truncation: responses.ResponseNewParamsTruncationAuto,
		Text: responses.ResponseTextConfigParam{
			Verbosity: responses.ResponseTextConfigVerbosityMedium,
		},
		Instructions: openai.String(systemPrompt),
		Input: responses.ResponseNewParamsInputUnion{
			OfString: openai.String(buildParaphrasePrompt(tone, text)),
		},
	}

	res, err := c.openaiClient.Responses.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("openapi responses api: %w", err)
	}

	return res.OutputText(), nil
}

const systemPrompt = `
You are a creative paraphrasing assistant.
Your task is to rephrase the given text to match the specified tone while preserving its original meaning.
Feel free to use synonyms but avoid uncommon words.
Use English in the paraphrased text.
Only respond with the paraphrased output.
`

func buildParaphrasePrompt(tone string, text string) string {
	return strings.TrimSpace(fmt.Sprintf(`
Tone: %s
Text: %s
	`, tone, text))
}
