package paraphrase_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gkatanacio/paraphraser-api/internal/paraphrase"
)

func Test_Payload_Unmarshal(t *testing.T) {
	data := []byte(`{"provider":"chatgpt","tone":"persuasive","text":"May I ask for a glass of water?"}`)

	var payload paraphrase.Payload
	err := json.Unmarshal(data, &payload)

	assert.NoError(t, err)
	assert.Equal(t, paraphrase.Payload{
		Provider: paraphrase.ChatGpt,
		Tone:     paraphrase.Persuasive,
		Text:     "May I ask for a glass of water?",
	}, payload)
}

func Test_Payload_Validate(t *testing.T) {
	testCases := map[string]struct {
		payload *paraphrase.Payload
		wantErr error
	}{
		"no error": {
			payload: &paraphrase.Payload{
				Provider: paraphrase.ChatGpt,
				Tone:     paraphrase.Formal,
				Text:     "Hey there! What's up?",
			},
			wantErr: nil,
		},
		"invalid provider": {
			payload: &paraphrase.Payload{
				Provider: "invalid",
				Tone:     paraphrase.Formal,
				Text:     "Hey there! What's up?",
			},
			wantErr: paraphrase.ErrInvalidProvider,
		},
		"invalid tone": {
			payload: &paraphrase.Payload{
				Provider: paraphrase.ChatGpt,
				Tone:     "invalid",
				Text:     "Hey there! What's up?",
			},
			wantErr: paraphrase.ErrInvalidTone,
		},
		"invalid text": {
			payload: &paraphrase.Payload{
				Provider: paraphrase.ChatGpt,
				Tone:     paraphrase.Formal,
				Text:     "",
			},
			wantErr: paraphrase.ErrInvalidText,
		},
	}

	for scenario, tc := range testCases {
		t.Run(scenario, func(t *testing.T) {
			assert.Equal(t, tc.wantErr, tc.payload.Validate())
		})
	}
}
