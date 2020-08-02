package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func postRequest() error {
	accessToken := os.Getenv("LINKEDIN_ACCESS_TOKEN")

	type shareCommentary struct {
		Text string `json:"text"`
	}

	type description struct {
		Text string `json:"text"`
	}

	type title struct {
		Text string `json:"text"`
	}

	type visibility struct {
		Kind string `json:"com.linkedin.ugc.MemberNetworkVisibility"`
	}

	type media struct {
		Status      string      `json:"status"`
		Description description `json:"description"`
		Media       string      `json:"media"`
		OriginalURL string      `json:"originalUrl"`
		Title       title       `json:"title"`
	}

	type content struct {
		ShareCommentary    shareCommentary `json:"shareCommentary"`
		ShareMediaCategory string          `json:"shareMediaCategory"`
		Media              []media         `json:"media"`
	}

	type specificContent struct {
		SpecificContent content `json:"com.linkedin.ugc.ShareContent"`
	}

	type body struct {
		Author          string          `json:"author"`
		LifecycleState  string          `json:"lifecycleState"`
		SpecificContent specificContent `json:"specificContent"`
		Visibility      visibility      `json:"visibility"`
	}

	// {
	// 	"author": "urn:li:person:8675309",
	// 	"lifecycleState": "PUBLISHED",
	// 	"specificContent": {
	// 		"com.linkedin.ugc.ShareContent": {
	// 			"shareCommentary": {
	// 				"text": "Learning more about LinkedIn by reading the LinkedIn Blog!"
	// 			},
	// 			"shareMediaCategory": "ARTICLE",
	// 			"media": [
	// 				{
	// 					"status": "READY",
	// 					"description": {
	// 						"text": "Official LinkedIn Blog - Your source for insights and information about LinkedIn."
	// 					},
	// 					"originalUrl": "https://blog.linkedin.com/",
	// 					"title": {
	// 						"text": "Official LinkedIn Blog"
	// 					}
	// 				}
	// 			]
	// 		}
	// 	},
	// 	"visibility": {
	// 		"com.linkedin.ugc.MemberNetworkVisibility": "CONNECTIONS"
	// 	}
	// }

	payload := body{
		Author:         "urn:li:person:poh_wxCNVI",
		LifecycleState: "PUBLISHED",
		SpecificContent: specificContent{
			SpecificContent: content{
				ShareCommentary: shareCommentary{
					Text: "Learning more about LinkedIn by reading the LinkedIn Blog!",
				},
				ShareMediaCategory: "ARTICLE",
				Media: []media{
					media{
						Status: "READY",
						Description: description{
							Text: "This is the description",
						},
						OriginalURL: "https://mihaiblebea.com",
						Title: title{
							Text: "This is the title",
						},
					},
				},
			},
		},
		Visibility: visibility{
			Kind: "PUBLIC",
		},
	}

	b, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	var jsonStr = []byte(`{
		"author": "urn:li:person:poh_wxCNVI",
		"lifecycleState": "PUBLISHED",
		"specificContent": {
			"com.linkedin.ugc.ShareContent": {
				"shareCommentary": {
					"text": "Learning more about LinkedIn by reading the LinkedIn Blog!"
				},
				"shareMediaCategory": "ARTICLE",
				"media": [
					{
						"status": "READY",
						"description": {
							"text": "Official LinkedIn Blog - Your source for insights and information about LinkedIn."
						},
						"originalUrl": "https://blog.linkedin.com/",
						"title": {
							"text": "Official LinkedIn Blog"
						}
					}
				]
			}
		},
		"visibility": {
			"com.linkedin.ugc.MemberNetworkVisibility": "CONNECTIONS"
		}
	}`)

	req, err := http.NewRequest("POST", "https://api.linkedin.com/v2/ugcPosts", bytes.NewBuffer(jsonStr))
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

	b, err = ioutil.ReadAll(resp.Body)
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
