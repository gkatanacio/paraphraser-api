package handler_test

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gkatanacio/paraphraser-api/internal/errs"
	"github.com/gkatanacio/paraphraser-api/internal/handler"
)

func Test_JsonResponse(t *testing.T) {
	testCases := map[string]struct {
		data       interface{}
		statusCode int
		wantBody   string
	}{
		"map object": {
			data: map[string]string{
				"firstName": "John",
				"lastName":  "Doe",
			},
			statusCode: 200,
			wantBody:   `{"firstName":"John","lastName":"Doe"}`,
		},
		"array": {
			data:       []bool{true, false, true},
			statusCode: 202,
			wantBody:   "[true,false,true]",
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			response, err := handler.JsonResponse(tc.data, tc.statusCode)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, response.Body)
			assert.Equal(t, tc.statusCode, response.StatusCode)
		})
	}
}

func Test_ErrorResponse(t *testing.T) {
	testCases := map[string]struct {
		err        error
		wantBody   string
		wantStatus int
	}{
		"generic error": {
			err:        errors.New("something went wrong"),
			wantBody:   `{"error":"something went wrong"}`,
			wantStatus: 500,
		},
		"bad request": {
			err:        errs.NewBadRequest("invalid input"),
			wantBody:   `{"error":"invalid input"}`,
			wantStatus: http.StatusBadRequest,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			response, err := handler.ErrorResponse(tc.err)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, response.Body)
			assert.Equal(t, tc.wantStatus, response.StatusCode)
		})
	}
}
