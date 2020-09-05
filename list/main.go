package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

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

	contacts, err := lService.GetContacts()
	if err != nil {
		logger.Error(err)
	}
	for _, con := range *contacts {
		fmt.Println(con.Email)
		in := sService.Build("mihaiserban.blebea@gmail.com", "<h1>This is an email</h1>", "Hello there!")
		out, err := sService.Send(in)
		if err != nil {
			logger.Error(err)
		}

		logger.Info(out)
	}
}
