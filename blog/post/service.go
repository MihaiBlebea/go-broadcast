package post

// Service interface
type Service interface {
	GetAllPosts() (*[]Post, error)
	BuildPost(filePath string) (*Post, error)
}
