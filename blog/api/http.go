package api

import (
	"net/http"
	"time"

	"github.com/MihaiBlebea/blog/go-broadcast/page"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type httpServer struct {
	pageService page.Service
	handler     http.Handler
	logger      *logrus.Logger
}

// HTTPServer interface
type HTTPServer interface {
	GetHandler() http.Handler
	BlogHandler(w http.ResponseWriter, r *http.Request)
	ArticleHandler(w http.ResponseWriter, r *http.Request)
}

// NewHTTPServer returns a new http server service
func NewHTTPServer(pageService page.Service, logger *logrus.Logger) HTTPServer {
	httpServer := httpServer{
		pageService: pageService,
		logger:      logger,
	}

	r := mux.NewRouter()

	r.Methods("GET").Path("/").HandlerFunc(httpServer.BlogHandler)
	r.Methods("GET").Path("/article/{slug}").HandlerFunc(httpServer.ArticleHandler)

	middlewareCORS := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		// Enable Debugging for testing, consider disabling in production
		Debug: false,
	})

	httpServer.handler = middlewareCORS.Handler(r)

	return &httpServer
}

func (h *httpServer) BlogHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.WithFields(logrus.Fields{
		"url": r.URL.String(),
	}).Info("HTTP request started")
	start := time.Now()

	defer h.logger.WithFields(logrus.Fields{
		"duration": time.Since(start).Nanoseconds(),
	}).Info("HTTP request ended")

	page, err := h.pageService.LoadBlogPage("/", nil)
	if err != nil {
		page, _ = h.pageService.LoadPage("/error", nil)
		page.Render(w)
		return
	}

	err = page.Render(w)
	if err != nil {
		page, _ = h.pageService.LoadPage("/error", nil)
		page.Render(w)
	}
}

func (h *httpServer) ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	slug := vars["slug"]

	params := struct {
		Name string
	}{
		Name: "Zlatan",
	}
	h.logger.WithFields(logrus.Fields{
		"url":    r.URL.String(),
		"params": params,
	}).Info("HTTP request started")
	start := time.Now()

	defer h.logger.WithFields(logrus.Fields{
		"duration": time.Since(start).Nanoseconds(),
	}).Info("HTTP request ended")

	page, err := h.pageService.LoadPage(slug, params)
	if err != nil {
		page, _ = h.pageService.LoadPage("/error", params)
		page.Render(w)
		return
	}

	err = page.Render(w)
	if err != nil {
		page, _ = h.pageService.LoadPage("/error", params)
		page.Render(w)
	}
}

func (h *httpServer) GetHandler() http.Handler {
	return h.handler
}
