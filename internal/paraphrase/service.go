package paraphrase

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrNoConfiguredParaphraser = errors.New("no paraphraser configured for provider")
)

// Paraphraser is the interface that will be implemented by concrete paraphrasing providers.
type Paraphraser interface {
	// Paraphrase rewords the given text to sound more like the given tone.
	Paraphrase(ctx context.Context, tone string, text string) (string, error)
}

// Service is the service layer containing functionality for paraphrasing.
type Service struct {
	cfg          *Config
	paraphrasers map[Provider]Paraphraser
}

func NewService(cfg *Config, paraphrasers map[Provider]Paraphraser) *Service {
	return &Service{
		cfg:          cfg,
		paraphrasers: paraphrasers,
	}
}

// Paraphrase forwards the paraphrasing request to the configured paraphraser for the given provider.
// A timeout is set for the request based on the service configuration.
func (s *Service) Paraphrase(ctx context.Context, provider Provider, tone Tone, text string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.cfg.Timeout)
	defer cancel()

	paraphraser, ok := s.paraphrasers[provider]
	if !ok {
		return "", ErrNoConfiguredParaphraser
	}

	result, err := paraphraser.Paraphrase(ctx, string(tone), text)
	if err != nil {
		return "", fmt.Errorf("paraphrase provider: %w", err)
	}

	return result, nil
}
