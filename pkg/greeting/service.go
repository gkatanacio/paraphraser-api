package greeting

import (
	"fmt"
)

type Servicer interface {
	HelloMessage(name string) string
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) HelloMessage(name string) string {
	return fmt.Sprintf("Hello there, %s!", name)
}
