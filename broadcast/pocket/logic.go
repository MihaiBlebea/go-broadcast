package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/sirupsen/logrus"
)

const baseURL = "https://getpocket.com/v3"

type service struct {
	consumerKey string
	accessToken string
	logger      *logrus.Logger
}

// New creates a new Pocket service
func New(consumerKey, accessToken string, logger *logrus.Logger) Service {
	return &service{
		consumerKey: consumerKey,
		accessToken: accessToken,
		logger:      logger,
	}
}

func (s *service) RetrieveArticles() (*RetrieveArticlesResponse, error) {
	url := fmt.Sprintf(
		"%s/%s",
		baseURL,
		"get",
	)

	payload := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
	}{
		ConsumerKey: s.consumerKey,
		AccessToken: s.accessToken,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return &RetrieveArticlesResponse{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &RetrieveArticlesResponse{}, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return &RetrieveArticlesResponse{}, err
	}

	return s.getArticles(body)
}

func (s *service) ArchiveArticle(itemID string) error {
	url := fmt.Sprintf(
		"%s/%s",
		baseURL,
		"send",
	)

	payload := struct {
		ConsumerKey string   `json:"consumer_key"`
		AccessToken string   `json:"access_token"`
		Actions     []Action `json:"actions"`
	}{
		ConsumerKey: s.consumerKey,
		AccessToken: s.accessToken,
		Actions: []Action{
			Action{ItemID: itemID, Action: "archive"},
		},
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) getArticles(body []byte) (*RetrieveArticlesResponse, error) {
	resp := new(RetrieveArticlesResponse)
	err := json.Unmarshal(body, &resp)
	if err != nil {
		return &RetrieveArticlesResponse{}, err
	}

	return resp, nil
}
