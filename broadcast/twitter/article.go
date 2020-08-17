package twitter

// Article interface
type Article interface {
	GetHashedTags() []string
	GetSummary() string
	GetTitle() string
	GetURL() string
}
