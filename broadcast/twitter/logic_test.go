package twitter

import (
	"fmt"
	"testing"

	"github.com/MihaiBlebea/broadcast/model"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func newTwitterClient() *service {
	config := oauth1.NewConfig("", "")
	tkn := oauth1.NewToken("", "")
	httpClient := config.Client(oauth1.NoContext, tkn)

	client := twitter.NewClient(httpClient)
	return &service{client, nil}
}

func genPlaceholderStr(count int) string {
	i := 0
	var summary string
	for i < count {
		summary += "a"
		i++
	}

	return summary
}

func TestCreateTrimmedSummary(t *testing.T) {
	article := model.Article{
		Title:   "Demo Article",
		URL:     "",
		Summary: genPlaceholderStr(1000),
		Tags:    []string{},
	}

	tests := []struct {
		In  int
		Out string
	}{
		{In: 2, Out: ""},
		{In: 3, Out: ""},
		{In: 4, Out: "a..."},
		{In: 5, Out: "aa..."},
	}

	twitter := newTwitterClient()

	for _, test := range tests {
		t.Run(test.Out, func(t *testing.T) {
			summary := twitter.createTrimmedSummary(&article, test.In)
			if len(summary) > test.In {
				t.Errorf("Expected summary to be %s, but got %s", test.Out, summary)
			}
		})
	}
}

func TestCreateTags(t *testing.T) {
	tags := []string{
		genPlaceholderStr(3),
		genPlaceholderStr(3),
		genPlaceholderStr(5),
	}

	article := model.Article{
		Title:   "Demo Article",
		URL:     "",
		Summary: "",
		Tags:    tags,
	}

	twitter := newTwitterClient()

	tagString := twitter.createTags(&article)

	expected := len(tags) - 1
	for _, t := range tags {
		expected += len(t) + 1
	}

	if len(tagString) != expected {
		t.Errorf("Expected tag string to be %d char long, but got %d char long", expected, len(tagString))
	}
}

func TestSummaryValidLength(t *testing.T) {
	article := model.Article{
		Title:   "Demo Article",
		URL:     genPlaceholderStr(2),
		Summary: genPlaceholderStr(1000),
		Tags: []string{
			genPlaceholderStr(2),
			genPlaceholderStr(2),
			genPlaceholderStr(2),
		},
	}

	twitter := newTwitterClient()

	content, err := twitter.createUpdateContent(&article)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(content)
	if len(content) > maxChar {
		t.Errorf("Expected content to be %d char long, but got %d char long", maxChar, len(content))
	}
}
