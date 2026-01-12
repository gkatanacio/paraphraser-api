package paraphrase_test

import (
	context "context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/gkatanacio/paraphraser-api/internal/paraphrase"

	mocksparaphrase "github.com/gkatanacio/paraphraser-api/mocks/paraphrase"
)

func Test_Service_Paraphrase_Success(t *testing.T) {
	testProvider := paraphrase.ChatGpt
	testTone := paraphrase.Amicable
	testText := "I'm hungry. What's for dinner?"
	testResult := "I have a strong appetite. May I know what we are having for dinner?"

	mockParaphraser := mocksparaphrase.NewMockParaphraser(t)
	mockParaphraser.EXPECT().Paraphrase(mock.Anything, string(testTone), testText).Return(testResult, nil).Once()

	service := paraphrase.NewService(
		&paraphrase.Config{
			Timeout: 3 * time.Second,
		},
		map[paraphrase.Provider]paraphrase.Paraphraser{
			testProvider: mockParaphraser,
		},
	)

	result, err := service.Paraphrase(context.Background(), testProvider, testTone, testText)

	assert.NoError(t, err)
	assert.Equal(t, testResult, result)
}

func Test_Service_Paraphrase_NoConfiguredParaphraser(t *testing.T) {
	service := paraphrase.NewService(
		&paraphrase.Config{
			Timeout: 3 * time.Second,
		},
		map[paraphrase.Provider]paraphrase.Paraphraser{},
	)

	_, err := service.Paraphrase(context.Background(), "unknown", paraphrase.Fun, "Hello world")

	assert.Equal(t, paraphrase.ErrNoConfiguredParaphraser, err)
}

func Test_Service_Paraphrase_ParaphraserError(t *testing.T) {
	testProvider := paraphrase.ChatGpt
	testErr := context.DeadlineExceeded

	mockParaphraser := mocksparaphrase.NewMockParaphraser(t)
	mockParaphraser.EXPECT().Paraphrase(mock.Anything, mock.Anything, mock.Anything).Return("", testErr).Once()

	service := paraphrase.NewService(
		&paraphrase.Config{
			Timeout: 1 * time.Millisecond,
		},
		map[paraphrase.Provider]paraphrase.Paraphraser{
			testProvider: mockParaphraser,
		},
	)

	_, err := service.Paraphrase(context.Background(), testProvider, paraphrase.Formal, "Hello world")

	assert.ErrorIs(t, err, testErr)
}
