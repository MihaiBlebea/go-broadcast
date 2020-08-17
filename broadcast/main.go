package main

import (
	"errors"
	"os"
	"time"

	"github.com/MihaiBlebea/broadcast/linkedin"
	"github.com/MihaiBlebea/broadcast/model"
	"github.com/MihaiBlebea/broadcast/pocket"
	"github.com/MihaiBlebea/broadcast/twitter"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
)

// errors
var (
	ErrNoArticle = errors.New("No articles left in Pocket")
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	pocket := pocket.New(
		os.Getenv("POCKET_CONSUMER_KEY"),
		os.Getenv("POCKET_ACCESS_TOKEN"),
		logger,
	)

	linkedin := linkedin.New(
		os.Getenv("LINKEDIN_ACCESS_TOKEN"),
		logger,
	)
	twitter := twitter.New(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
		os.Getenv("TWITTER_TOKEN"),
		os.Getenv("TWITTER_TOKEN_SECRET"),
		logger,
	)

	c := cron.New()
	c.AddFunc("0 0 16 * * *", func() {
		logger.Info("Cronjob running - broadcasting article")
		err := publish(pocket, linkedin, twitter)
		if err != nil {
			logger.Error(err)
		}
	})

	c.Start()

	for true {
		logger.Info("Starting the script. Daily check")
		time.Sleep(24 * 60 * time.Minute)
	}
}

func publish(pocket pocket.Service, linkedin linkedin.Service, twitter twitter.Service) error {
	resp, err := pocket.RetrieveArticles()
	if err != nil {
		return err
	}

	if len(resp.GetArticles()) == 0 {
		return ErrNoArticle
	}

	article := resp.GetArticles()[0]

	publish := model.Article{
		Title:   article.ResolvedTitle,
		URL:     article.GivenURL,
		Summary: article.Excerpt,
		Tags:    article.GetTags(),
	}

	err = linkedin.ShareArticle(&publish)
	if err != nil {
		return err
	}

	err = twitter.PostTweet(&publish)
	if err != nil {
		return err
	}

	err = pocket.ArchiveArticle(article.ItemID)

	return err
}
