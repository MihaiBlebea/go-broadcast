package linkedin

// Service interface
type Service interface {
	ShareArticle(article Article) error
}
