package twitter

import "github.com/MihaiBlebea/broadcast/model"

// Service interface
type Service interface {
	PostTweet(article *model.Article) error
}
