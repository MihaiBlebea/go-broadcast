package main

import (
	"context"
	"encoding/json"
	"fmt"

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

	credentials := `{
		"web": {
			"client_id": "326310223601-o53hnchgeuojbcn3s3k9mo68iqh3i4aj.apps.googleusercontent.com",
			"project_id": "go-list-1599302644444",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_secret": "TLCsgg8EHYoWqZHUrTNZRrdX",
			"redirect_uris": [
				"https://mihaiblebea.com"
			]
		}
	}`

	b := []byte(credentials)

	config, err := google.ConfigFromJSON(b, people.ContactsReadonlyScope)
	if err != nil {
		logger.Error(err)
	}
	token := `{
		"access_token": "ya29.a0AfH6SMDy_TgA4PCOgT5ZnuLxMnCFCc5qyE9wzlCc0jPE6SrnQda_sXdMR59LsBSYdYobS-MUS6pTYzfgkpVQq32aO5wvZwgOOlHSUod7A2zKrkXLTImLyKqE-vaL7iatd536az0XkLeSG0mnWJzYSx1zwEomWtjrfgQ",
		"token_type": "Bearer",
		"refresh_token": "1//03fnM5tLmq0wgCgYIARAAGAMSNwF-L9IrqUpybmgzOlcvDh8OUnOoFdn13-JNzrePQsL-S_xSa8OBD0CgXFHbmG0B0N0axg4t1io",
		"expiry": "2020-09-05T12:56:59.638099+01:00"
	}`
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

	sService, err := sender.New("eu-west-2", "mihai@mihaiblebea.com")
	if err != nil {
		logger.Error(err)
	}

	contacts, err := lService.GetContacts()
	if err != nil {
		logger.Error(err)
	}
	for _, con := range *contacts {
		fmt.Println(con.Email)
		in := sService.Build(con.Email, "<h1>This is an email</h1>", "Hello there!")
		out, err := sService.Send(in)
		if err != nil {
			logger.Error(err)
		}

		logger.Info(out)
	}
}
