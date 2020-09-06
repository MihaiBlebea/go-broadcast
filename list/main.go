package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/MihaiBlebea/list/api"
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

	// credentials := os.Getenv("GOOGLE_CREDENTIALS_FILE")
	credentials := `{
		"web": {
			"client_id": "724740853057-ammjchsnvludb04mha146rdpqk8nc8le.apps.googleusercontent.com",
			"project_id": "go-list-1599300866421",
			"auth_uri": "https://accounts.google.com/o/oauth2/auth",
			"token_uri": "https://oauth2.googleapis.com/token",
			"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
			"client_secret": "eVLCDSe01bzECAiyvJBt21mJ",
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

	// token := os.Getenv("GOOGLE_TOKEN_FILE")
	token := `{
		"access_token": "ya29.a0AfH6SMDyIhL1v4TRXnyrphyz5azIzytkn_-PVqy-GzrIASiyyFivZ1sDLm8lasZTWy4jkhZKqitdfDZNyUhwil_MA7Dgyoyxj35_V5M8T1x0_1hsSpSU3zj_vchHJ8ArQxhaOSCZvdIF4cby8STIKQ4VcFqvq3wdK80",
		"token_type": "Bearer",
		"refresh_token": "1//03DaM3MRRor8-CgYIARAAGAMSNwF-L9Irh7CqCoXyG3zeuWyB1qvKIq8qeQsIpwWGnGUixfqaeqEBTSI0Bz56EMBy9RKCDVx95_U",
		"expiry": "2020-09-06T18:50:53.096516+01:00"
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

	sService, err := sender.New("eu-west-2", "mihaiserban.blebea@gmail.com")
	if err != nil {
		logger.Error(err)
	}

	cService := core.New(lService, sService, logger)

	httpServer := api.New(lService, logger)

	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	logger.Info(fmt.Sprintf("Application started HTTP server on port %s", httpPort))

	go func() {
		err = http.ListenAndServe(httpPort, *httpServer.Handler())
		if err != nil {
			logger.Fatal(err)
		}
	}()

	for {
		logger.Info("Running the loop")

		fmt.Println(cService)
		// err = cService.Execute()
		// if err != nil {
		// 	logger.Error(err)
		// }

		time.Sleep(time.Second * 60 * 60)
	}
}
