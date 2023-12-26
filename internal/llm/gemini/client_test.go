package gemini_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gkatanacio/paraphraser-api/internal/llm/gemini"
	"github.com/gkatanacio/paraphraser-api/internal/testutil"
)

func Test_Client_Paraphrase(t *testing.T) {
	testutil.IntegrationTest(t)

	client := gemini.NewClient(gemini.ConfigFromEnv())
	result, err := client.Paraphrase(context.Background(), "formal", "I'm hungry. What's for dinner?")

	assert.NoError(t, err)
	assert.NotEmpty(t, result)
}
