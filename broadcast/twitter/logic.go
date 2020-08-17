package twitter

import (
	"errors"
	"fmt"
	"strings"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/sirupsen/logrus"
)

// Max number of chars that a twitter update can have.
// Can change yearly
const maxChar = 280

// Errors
var (
	ErrInvalidContentLength = errors.New("Content is too long to be posted on Twitter")
)

type service struct {
	client *twitter.Client
	logger *logrus.Logger
}

// New returns a new twitter service
func New(consumerKey, consumerSecret, token, tokenSecret string, logger *logrus.Logger) Service {
	config := oauth1.NewConfig(
		consumerKey,
		consumerSecret,
	)
	tkn := oauth1.NewToken(
		token,
		tokenSecret,
	)
	httpClient := config.Client(oauth1.NoContext, tkn)

	// Twitter client
	client := twitter.NewClient(httpClient)
	return &service{client, logger}
}

func (s *service) PostTweet(article Article) error {
	// Send a Tweet
	content, err := s.createUpdateContent(article)
	if err != nil {
		return err
	}
	_, _, err = s.client.Statuses.Update(content, nil)

	return err
}

func (s *service) createUpdateContent(article Article) (string, error) {
	currentCount := maxChar
	tags := s.createTags(article)

	currentCount -= len(article.GetURL())
	currentCount -= len(tags)
	currentCount -= 2 // 2 breaking line characters
	if currentCount <= 0 {
		return "", ErrInvalidContentLength
	}

	trimmedSummary := article.GetSummary()
	if len(article.GetSummary()) >= currentCount {
		trimmedSummary = s.createTrimmedSummary(article, currentCount)
	}

	return fmt.Sprintf(
		"%s\n%s\n%s",
		trimmedSummary,
		tags,
		article.GetURL(),
	), nil
}

func (s *service) createTags(article Article) string {
	if len(article.GetHashedTags()) == 0 {
		return ""
	}

	return strings.Join(article.GetHashedTags(), " ")
}

func (s *service) createTrimmedSummary(article Article, limit int) string {
	if limit-3 <= 0 {
		return ""
	}

	return fmt.Sprintf(
		"%s...",
		article.GetSummary()[0:limit-3], // take in consideration the ...
	)
}
