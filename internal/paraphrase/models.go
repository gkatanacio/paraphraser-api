package paraphrase

import (
	"errors"
	"slices"
)

// Provider is the type for representing possible paraphrasing providers.
type Provider string

const (
	ChatGpt Provider = "chatgpt"
	Bard    Provider = "bard"
	Grok    Provider = "grok"
)

// var availableProviders = []Provider{ChatGpt, Bard, Grok}
var availableProviders = []Provider{ChatGpt} // only ChatGpt is available for now

// Tone is the type for representing possible paraphrasing tones.
type Tone string

const (
	Formal      Tone = "formal"
	Amicable    Tone = "amicable"
	Fun         Tone = "fun"
	Casual      Tone = "casual"
	Sympathetic Tone = "sympathetic"
	Persuasive  Tone = "persuasive"
)

var availableTones = []Tone{Formal, Amicable, Fun, Casual, Sympathetic, Persuasive}

var (
	ErrInvalidProvider = errors.New("invalid value for provider")
	ErrInvalidTone     = errors.New("invalid value for tone")
	ErrInvalidText     = errors.New("invalid value for text")
)

// Payload represents the body for the paraphrase request.
type Payload struct {
	Provider `json:"provider"`
	Tone     `json:"tone"`
	Text     string `json:"text"`
}

// Validate throws an error if the Payload is invalid.
func (p *Payload) Validate() error {
	if slices.Index(availableProviders, p.Provider) == -1 {
		return ErrInvalidProvider
	}

	if slices.Index(availableTones, p.Tone) == -1 {
		return ErrInvalidTone
	}

	if len(p.Text) == 0 {
		return ErrInvalidText
	}

	return nil
}
