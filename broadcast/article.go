package main

// Article model
type Article struct {
	Image     string   `json:"image"`
	Title     string   `json:"title"`
	Slug      string   `json:"slug"`
	Summary   string   `json:"summary"`
	Tags      []string `json:"tags"`
	Published string   `json:"published"`
}
