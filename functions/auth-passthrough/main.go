package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handle(
	ctx context.Context,
	request *events.APIGatewayV2CustomAuthorizerV2Request,
) (*events.APIGatewayV2CustomAuthorizerSimpleResponse, error) {
	return &events.APIGatewayV2CustomAuthorizerSimpleResponse{
		IsAuthorized: true,
	}, nil
}

func main() {
	lambda.Start(handle)
}
