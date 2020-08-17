package model

import "fmt"

// Article model
type Article struct {
	Title   string   `json:"title"`
	URL     string   `json:"url"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}

// GetTitle returns the article title
func (a *Article) GetTitle() string {
	return a.Title
}

// GetURL returns the article url
func (a *Article) GetURL() string {
	return a.URL
}

// GetSummary returns the article summary
func (a *Article) GetSummary() string {
	return a.Summary
}

// GetHashedTags returns the article title
func (a *Article) GetHashedTags() []string {
	tags := []string{}
	for _, tag := range a.Tags {
		tags = append(tags, fmt.Sprintf("#%s", tag))
	}

	return tags
}
