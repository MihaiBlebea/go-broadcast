package pocket

import (
	"time"
)

// Article struct
type Article struct {
	ItemID        string `json:"item_id"`
	GivenURL      string `json:"given_url"`
	ResolvedTitle string `json:"resolved_title"`
	Excerpt       string `json:"excerpt"`
}

// RetrieveArticlesResponse response from retrieve endpoint
type RetrieveArticlesResponse struct {
	Status int                `json:"status"`
	List   map[string]Article `json:"list"`
	Error  string             `json:"error"`
	Since  int64              `json:"since"`
}

// GetSinceTimestamp returns the time.Time of the string Since attribute
func (r *RetrieveArticlesResponse) GetSinceTimestamp() (time.Time, error) {
	return time.Unix(r.Since, 0), nil
}

// GetArticles returns all articles in the list
func (r *RetrieveArticlesResponse) GetArticles() []Article {
	var articles []Article
	for _, article := range r.List {
		articles = append(articles, article)
	}

	return articles
}

// Action for the modify article endpoint
type Action struct {
	Action string `json:"action"`
	ItemID string `json:"item_id"`
}
