package linkedin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/MihaiBlebea/broadcast/model"
	"github.com/sirupsen/logrus"
)

const baseURL = "https://api.linkedin.com/v2"

type service struct {
	accessToken string
	logger      *logrus.Logger
}

// New creates a new linkedin service
func New(accessToken string, logger *logrus.Logger) Service {
	return &service{accessToken, logger}
}

func (s *service) ShareArticle(article *model.Article) error {
	payload, err := s.createNewPayload(article)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/%s", baseURL, "ugcPosts"), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Restli-Protocol-Version", "2.0.0")
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", s.accessToken),
	)

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

	if resp.StatusCode != http.StatusCreated {
		return errors.New("Something went wrong with the server")
	}

	return nil
}

func (s *service) createNewPayload(article *model.Article) ([]byte, error) {
	payload := make(map[string]interface{})

	shareCommentary := make(map[string]string)
	shareCommentary["text"] = article.Title

	description := make(map[string]string)
	description["text"] = article.Summary

	title := make(map[string]string)
	title["text"] = article.Title

	media := make(map[string]interface{})
	media["status"] = "READY"
	media["description"] = description
	media["originalUrl"] = article.URL
	media["title"] = title

	shareContent := make(map[string]interface{})
	shareContent["shareCommentary"] = shareCommentary
	shareContent["shareMediaCategory"] = "ARTICLE"
	shareContent["media"] = []map[string]interface{}{media}

	specificContent := make(map[string]interface{})
	specificContent["com.linkedin.ugc.ShareContent"] = shareContent

	visibility := make(map[string]string)
	visibility["com.linkedin.ugc.MemberNetworkVisibility"] = "CONNECTIONS"

	payload["author"] = "urn:li:person:poh_wxCNVI"
	payload["lifecycleState"] = "PUBLISHED"
	payload["specificContent"] = specificContent
	payload["visibility"] = visibility

	p, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	return p, nil
}
