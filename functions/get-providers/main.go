package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gkatanacio/paraphraser-api/internal/handler"
	"github.com/gkatanacio/paraphraser-api/internal/paraphrase"
)

func handle(ctx context.Context) (*events.APIGatewayV2HTTPResponse, error) {
	return handler.JsonResponse(http.StatusOK, map[string][]paraphrase.Provider{"providers": paraphrase.AvailableProviders})
}

func main() {
	lambda.Start(handle)
}
