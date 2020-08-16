package pocket

// Service to interact with pocket.com api
type Service interface {
	RetrieveArticles() (*RetrieveArticlesResponse, error)
	ArchiveArticle(itemID string) error
}
