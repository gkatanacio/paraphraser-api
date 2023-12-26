package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gkatanacio/paraphraser-api/internal/errs"
	"github.com/gkatanacio/paraphraser-api/internal/handler"
	"github.com/gkatanacio/paraphraser-api/internal/llm/chatgpt"
	"github.com/gkatanacio/paraphraser-api/internal/llm/gemini"
	"github.com/gkatanacio/paraphraser-api/internal/paraphrase"
)

var paraphraseService *paraphrase.Service

func init() {
	paraphraseService = paraphrase.NewService(
		paraphrase.ConfigFromEnv(),
		map[paraphrase.Provider]paraphrase.Paraphraser{
			paraphrase.ChatGpt: chatgpt.NewClient(chatgpt.ConfigFromEnv()),
			paraphrase.Gemini:  gemini.NewClient(gemini.ConfigFromEnv()),
		},
	)
}

func handle(ctx context.Context, request *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	slog.Info(fmt.Sprintf("request: %+v", request))

	var payload paraphrase.Payload
	if err := json.Unmarshal([]byte(request.Body), &payload); err != nil {
		slog.Error(err.Error())
		return handler.ErrorResponse(errs.NewBadRequest("malformed request body"))
	}

	if err := payload.Validate(); err != nil {
		slog.Error(err.Error())
		return handler.ErrorResponse(errs.NewBadRequest(err.Error()))
	}

	result, err := paraphraseService.Paraphrase(ctx, payload.Provider, payload.Tone, payload.Text)
	if err != nil {
		slog.Error(err.Error())
		return handler.ErrorResponse(errors.New("failed to paraphrase text"))
	}

	return handler.JsonResponse(map[string]string{"result": result}, 200)
}

func main() {
	lambda.Start(handle)
}
