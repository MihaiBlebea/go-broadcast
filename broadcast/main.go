package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	article := Article{
		Image:     "https://demo.plugins360.com/wp-content/uploads/2017/12/demo.png",
		Title:     "This is a demo title",
		Slug:      "demo-slug",
		Summary:   "This is a short description of the internal content of the blog article",
		Tags:      []string{"article", "it", "tech"},
		Published: "2020-07-28",
	}

	linkedin := linkedinService{logger}
	err := linkedin.ShareArticle(&article)
	if err != nil {
		logger.Error(err)
	}

	twitter := twitterService{logger}
	err = twitter.PostTweet(&article)
	if err != nil {
		logger.Error(err)
	}
}
