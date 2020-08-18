package api

import (
	"fmt"
	"net/http"
	"os"
	"path"
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
	r.Methods("GET").Path("/contact").HandlerFunc(httpServer.GetContactHandler)

	r.Methods("GET").Path("/about").HandlerFunc(httpServer.GetAboutHandler)

	r.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir(
					httpServer.staticFolderPath(),
				),
			),
		),
	)

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

	page, err := h.pageService.LoadArticlePage(slug, params)
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

func (h *httpServer) GetContactHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.WithFields(logrus.Fields{
		"url": r.URL.String(),
	}).Info("HTTP request started")
	start := time.Now()

	defer h.logger.WithFields(logrus.Fields{
		"duration": time.Since(start).Nanoseconds(),
	}).Info("HTTP request ended")

	page, err := h.pageService.LoadPage("contact", nil)
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

func (h *httpServer) GetAboutHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.WithFields(logrus.Fields{
		"url": r.URL.String(),
	}).Info("HTTP request started")
	start := time.Now()

	defer h.logger.WithFields(logrus.Fields{
		"duration": time.Since(start).Nanoseconds(),
	}).Info("HTTP request ended")

	page, err := h.pageService.LoadPage("about", nil)
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

func (h *httpServer) GetHandler() http.Handler {
	return h.handler
}

func (h *httpServer) staticFolderPath() string {
	p, err := os.Executable()
	if err != nil {
		h.logger.Fatal(err)
	}

	absPath := fmt.Sprintf(
		"%s/%s/",
		path.Dir(p),
		"static",
	)
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		h.logger.Fatal(err)
	}

	return absPath
}
