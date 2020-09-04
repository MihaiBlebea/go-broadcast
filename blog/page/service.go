package page

// Service interface for page package
type Service interface {
	LoadTemplate(URL string) (*Page, error)
	LoadStaticFile(URL string) ([]byte, error)
}
