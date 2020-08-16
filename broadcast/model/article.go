package model

// Article model
type Article struct {
	Title   string   `json:"title"`
	URL     string   `json:"url"`
	Summary string   `json:"summary"`
	Tags    []string `json:"tags"`
}
