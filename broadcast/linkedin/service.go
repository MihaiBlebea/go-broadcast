package linkedin

import "github.com/MihaiBlebea/broadcast/model"

// Service interface
type Service interface {
	ShareArticle(article *model.Article) error
}
