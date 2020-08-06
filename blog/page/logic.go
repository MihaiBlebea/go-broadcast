package page

import (
	"fmt"
	"html/template"
	"sort"

	"github.com/MihaiBlebea/blog/go-broadcast/assets"
	"github.com/MihaiBlebea/blog/go-broadcast/cache"
	"github.com/sirupsen/logrus"
)

type service struct {
	markdownService Markdown
	cache           *cache.Cache
	logger          *logrus.Logger
}

// New returns a new page service
func New(markdownService Markdown, cache *cache.Cache, logger *logrus.Logger) Service {
	return &service{
		markdownService: markdownService,
		cache:           cache,
		logger:          logger,
	}
}

func (s *service) getPartials(folder string) (string, error) {
	partials, err := assets.AssetDir(folder)
	if err != nil {
		return "", err
	}

	var result string
	for _, partial := range partials {
		b, err := assets.Asset(
			fmt.Sprintf("%s/%s", folder, partial),
		)
		if err != nil {
			return "", err
		}

		result += string(b)
	}

	return result, nil
}

func (s *service) LoadPage(slug string, optionalParams interface{}) (*Page, error) {
	p, err := s.cache.FindPage(slug)
	if err != nil {
		return &Page{}, err
	}
	page := p.(Page)

	lb, err := assets.Asset(
		fmt.Sprintf("templates/%s.gohtml", page.Layout),
	)
	if err != nil {
		return &Page{}, err
	}

	partials, err := s.getPartials("templates/partials")
	if err != nil {
		return &Page{}, err
	}

	tmpl, err := template.New("Template").Parse(
		partials + string(lb),
	)
	if err != nil {
		return &Page{}, err
	}

	page.Params = optionalParams
	page.Template = tmpl

	return &page, nil
}

func (s *service) LoadBlogPage(slug string, optionalParams interface{}) (*Page, error) {
	page, err := s.LoadPage(slug, optionalParams)
	if err != nil {
		return &Page{}, err
	}

	for _, p := range s.cache.All() {
		article := p.(Page)
		if article.Kind == kindArticle {
			page.Articles = append(page.Articles, &article)
		}
	}

	sort.SliceStable(page.Articles, func(i, j int) bool {
		return page.Articles[i].Published.After(page.Articles[j].Published)
	})

	return page, err
}

func (s *service) LoadArticlePage(slug string, optionalParams interface{}) (*Page, error) {
	page, err := s.LoadPage(slug, optionalParams)
	if err != nil {
		return &Page{}, err
	}

	for _, p := range s.cache.All() {
		article := p.(Page)

		if article.Slug == page.Slug || article.Kind != kindArticle {
			continue
		}

		for _, tag := range page.Tags {
			if contains(article.Tags, tag) {
				page.Articles = append(page.Articles, &article)
				break
			}
		}
	}

	return page, err
}

func contains(list []string, needle string) bool {
	for _, item := range list {
		if item == needle {
			return true
		}
	}

	return false
}
