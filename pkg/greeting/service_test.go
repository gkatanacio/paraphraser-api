package greeting_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gkatanacio/go-serverless-template/pkg/greeting"
)

func Test_Service_HelloMessage(t *testing.T) {
	// given
	service := greeting.NewService()
	name := "John Doe"

	// when
	msg := service.HelloMessage(name)

	// then
	assert.Equal(t, "Hello there, John Doe!", msg)
}
