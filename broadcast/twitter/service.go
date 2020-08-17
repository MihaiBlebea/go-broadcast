package twitter

// Service interface
type Service interface {
	PostTweet(article Article) error
}
