package page

// Service interface for page package
type Service interface {
	LoadPage(slug string, optionalParams interface{}) (*Page, error)
	LoadBlogPage(slug string, optionalParams interface{}) (*Page, error)
	LoadArticlePage(slug string, optionalParams interface{}) (*Page, error)
}
