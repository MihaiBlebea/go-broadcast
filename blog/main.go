package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/MihaiBlebea/blog/go-broadcast/api"
	"github.com/MihaiBlebea/blog/go-broadcast/cache"
	"github.com/MihaiBlebea/blog/go-broadcast/page"
	"github.com/MihaiBlebea/blog/go-broadcast/post"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	cache := cache.New()
	postService := post.New()
	pageService := page.New(postService, cache, logger)

	server := api.NewHTTPServer(pageService, logger)

	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	logger.Info(fmt.Sprintf("Application started HTTP server on port %s", httpPort))

	err := http.ListenAndServe(httpPort, server.GetHandler())
	if err != nil {
		logger.Fatal(err)
	}
}
