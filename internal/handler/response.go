package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aws/aws-lambda-go/events"

	"github.com/gkatanacio/paraphraser-api/internal/errs"
)

// JsonResponse builds the API Gateway Lambda response in JSON format.
func JsonResponse(data interface{}, statusCode int) (*events.APIGatewayV2HTTPResponse, error) {
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: statusCode,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(b),
	}, nil
}

type errorResponseBody struct {
	Error string `json:"error"`
}

// ErrorResponse builds the API Gateway Lambda response as a structured JSON
// for a given error. The HTTP status code is inferred from the error type.
func ErrorResponse(err error) (*events.APIGatewayV2HTTPResponse, error) {
	body := &errorResponseBody{}
	body.Error = err.Error()

	var status int
	var httpErr errs.HttpError
	switch {
	case errors.As(err, &httpErr):
		status = httpErr.StatusCode()
	default:
		status = http.StatusInternalServerError
	}

	return JsonResponse(body, status)
}
