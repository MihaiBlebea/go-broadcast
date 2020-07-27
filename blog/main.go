package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MihaiBlebea/blog/go-broadcast/api"
	"github.com/MihaiBlebea/blog/go-broadcast/assets"
	"github.com/MihaiBlebea/blog/go-broadcast/cache"
	"github.com/MihaiBlebea/blog/go-broadcast/page"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	markdown := page.NewMarkdown()
	cache := cache.New()

	pages, err := readAllMarkdownFiles(markdown)
	if err != nil {
		logger.Fatal(err)
	}

	for _, page := range pages {
		cache.StorePage(page.Slug, page)
	}

	service := page.New(markdown, cache, logger)

	server := api.NewHTTPServer(service, logger)

	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	logger.Info(fmt.Sprintf("Application started HTTP server on port %s", httpPort))

	err = http.ListenAndServe(httpPort, server.GetHandler())
	if err != nil {
		logger.Fatal(err)
	}
}

func readAllMarkdownFiles(markdownService page.Markdown) ([]page.Page, error) {
	files, err := assets.AssetDir("markdown")
	if err != nil {
		return []page.Page{}, err
	}

	var pages []page.Page

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
