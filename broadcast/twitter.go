package main

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func postTwitter() error {
	config := oauth1.NewConfig(
		os.Getenv("TWITTER_CONSUMER_KEY"),
		os.Getenv("TWITTER_CONSUMER_SECRET"),
	)
	token := oauth1.NewToken(
		os.Getenv("TWITTER_TOKEN"),
		os.Getenv("TWITTER_TOKEN_SECRET"),
	)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)

	// Home Timeline
	// tweets, resp, err := client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
	// 	Count: 20,
	// })

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update("just setting up my twttr", nil)
	if err != nil {
		return err
	}
	fmt.Println("Tweet", tweet.Text)
	fmt.Println("Response", resp)
	// Status Show
	// tweet, resp, err := client.Statuses.Show(585613041028431872, nil)

	// Search Tweets
	// search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
	// 	Query: "gopher",
	// })

	// User Show
	// user, resp, err := client.Users.Show(&twitter.UserShowParams{
	// 	ScreenName: "dghubble",
	// })

	// Followers
	// followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})

	return nil
}
