package hello

import (
	"errors"
	"fmt"
	"strings"
)

// Errors
var (
	ErrNoName     = errors.New("Name parameter is invalid")
	ErrNoAge      = errors.New("Age parameter is invalid")
	ErrNoTemplate = errors.New("Template parameter is invalid")
)

type service struct {
}

// New returns a new hello world service
func New() Service {
	return &service{}
}

func (s *service) BroadcastMessage(name string, age int, template string) (string, error) {
	if err := s.validateInput(name, age, template); err != nil {
		return "", err
	}

	return fmt.Sprintf(template, name, age), nil
}

func (s *service) validateInput(name string, age int, template string) error {
	if name == "" {
		return ErrNoName
	}

	if age == 0 {
		return ErrNoAge
	}

	if template == "" || strings.Contains(template, "%d") == false || strings.Contains(template, "%s") == false {
		return ErrNoTemplate
	}

	return nil
}
