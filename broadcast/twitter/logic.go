package twitter

import (
	"fmt"

	"github.com/MihaiBlebea/broadcast/model"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/sirupsen/logrus"
)

// Max number of chars that a twitter update can have.
// Can change yearly
const maxChar = 280

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

func (s *service) PostTweet(article *model.Article) error {
	// Send a Tweet
	_, _, err := s.client.Statuses.Update(s.createUpdateContent(article), nil)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) createUpdateContent(article *model.Article) string {
	currentCount := maxChar

	currentCount -= len(article.URL)
	currentCount -= 5

	trimmedSummary := article.Summary
	if len(article.Summary) >= currentCount {
		trimmedSummary = article.Summary[0:currentCount]
	}

	return fmt.Sprintf(
		"%s... \n %s",
		trimmedSummary,
		article.URL,
	)
}
