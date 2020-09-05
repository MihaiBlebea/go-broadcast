package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	"github.com/MihaiBlebea/list/core"
	"github.com/MihaiBlebea/list/sender"

	"github.com/MihaiBlebea/list/lists"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/people/v1"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	credentials := os.Getenv("GOOGLE_CREDENTIALS_FILE")

	b := []byte(credentials)

	config, err := google.ConfigFromJSON(b, people.ContactsReadonlyScope)
	if err != nil {
		logger.Error(err)
	}

	token := os.Getenv("GOOGLE_TOKEN_FILE")
	b = []byte(token)

	tok := &oauth2.Token{}
	err = json.Unmarshal(b, tok)
	if err != nil {
		logger.Error(err)
	}

	client := config.Client(context.Background(), tok)

	lService, err := lists.New(client)
	if err != nil {
		logger.Error(err)
	}

	sService, err := sender.New("eu-west-2", "mihaiserban.blebea@gmail.com")
	if err != nil {
		logger.Error(err)
	}

	cService := core.New(lService, sService, logger)

	for {
		logger.Info("Running the loop")

		err = cService.Execute()
		if err != nil {
			logger.Error(err)
		}

		time.Sleep(time.Second * 60 * 60)
	}
}
