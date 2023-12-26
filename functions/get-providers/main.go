package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gkatanacio/paraphraser-api/internal/handler"
	"github.com/gkatanacio/paraphraser-api/internal/paraphrase"
)

func handle(ctx context.Context) (*events.APIGatewayV2HTTPResponse, error) {
	return handler.JsonResponse(map[string][]paraphrase.Provider{"providers": paraphrase.AvailableProviders}, 200)
}

func main() {
	lambda.Start(handle)
}
