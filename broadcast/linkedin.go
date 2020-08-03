package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

type linkedinService struct {
	logger *logrus.Logger
}

// LinkedinService interface
type LinkedinService interface {
	ShareArticle(article *Article) error
}

func (s *linkedinService) ShareArticle(article *Article) error {
	accessToken := os.Getenv("LINKEDIN_ACCESS_TOKEN")

	payload, err := s.createNewPayload(article)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.linkedin.com/v2/ugcPosts", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Restli-Protocol-Version", "2.0.0")
	req.Header.Set(
		"Authorization",
		fmt.Sprintf("Bearer %s", accessToken),
	)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("Status code is %d", resp.StatusCode))
	fmt.Println(string(b))

	if resp.StatusCode != http.StatusCreated {
		return errors.New("Something went wrong with the server")
	}

	return nil
}

func (s *linkedinService) createNewPayload(article *Article) ([]byte, error) {
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
	media["originalUrl"] = "https://mihaiblebea.com/article/" + article.Slug
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

	fmt.Println(string(p))

	return p, nil
}
