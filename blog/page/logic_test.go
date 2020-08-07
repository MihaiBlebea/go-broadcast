package page

import (
	"fmt"
	"testing"

	"github.com/MihaiBlebea/blog/go-broadcast/assets"
	"github.com/MihaiBlebea/blog/go-broadcast/cache"
	"github.com/sirupsen/logrus"
)

func readAllMarkdownFiles(markdownService Markdown) ([]Page, error) {
	files, err := assets.AssetDir("markdown")
	if err != nil {
		return []Page{}, err
	}

	var pages []Page

	for _, fileName := range files {
		content, err := assets.Asset(
			fmt.Sprintf("markdown/%s", fileName),
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		p, err := markdownService.BuildPage(content)
		if err != nil {
			fmt.Println(err)
			continue
		}

		pages = append(pages, *p)
	}

	return pages, nil
}

func createService() (Service, error) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	markdown := NewMarkdown()
	cache := cache.New()

	pages, err := readAllMarkdownFiles(markdown)
	if err != nil {
		return &service{}, err
	}

	for _, page := range pages {
		cache.StorePage(page.Slug, page)
	}

	return New(markdown, cache, logger), nil
}

func TestGetValidPageBySlug(t *testing.T) {
	slug := "first-article"

	service, err := createService()
	if err != nil {
		t.Error(err)
		return
	}

	page, err := service.LoadArticlePage(slug, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if page.Slug != slug {
		t.Errorf("Expected slug %s, got %s", slug, page.Slug)
	}
}
