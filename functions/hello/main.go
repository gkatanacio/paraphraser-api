package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/gkatanacio/go-serverless-template/pkg/greeting"
)

var greetingService greeting.Servicer

func init() {
	greetingService = greeting.NewService()
}

func handler(ctx context.Context, event events.CloudWatchEvent) error {
	msg := greetingService.HelloMessage(os.Getenv("GREET_WHO"))
	log.Print(msg)
	return nil
}

func main() {
	lambda.Start(handler)
}
