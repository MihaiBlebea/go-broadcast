package leads

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

// Errors _
var (
	ErrRequestIssue = errors.New("Something went wrong with the request")
)

type service struct {
	url string
}

// New returns a new leads service
func New(url string) Service {
	return &service{url}
}

func (s *service) Save(name, email string) error {
	body := struct {
		Name  string
		Email string
	}{
		Name:  name,
		Email: email,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(b))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return ErrRequestIssue
	}

	return nil
}

func (s *service) FormatName(name string) string {
	return strings.Title(name)
}

func (s *service) FormatEmail(email string) string {
	return strings.ToLower(email)
}
