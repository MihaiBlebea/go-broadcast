package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/MihaiBlebea/blog/go-broadcast/page"
	"github.com/gorilla/mux"
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
	TemplateHandler(w http.ResponseWriter, r *http.Request)
}

// NewHTTPServer returns a new http server service
func NewHTTPServer(pageService page.Service, logger *logrus.Logger) HTTPServer {
	httpServer := httpServer{
		pageService: pageService,
		logger:      logger,
	}

	r := mux.NewRouter()

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

	r.PathPrefix("/").HandlerFunc(httpServer.TemplateHandler)

	httpServer.handler = r

	return &httpServer
}

func (h *httpServer) TemplateHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	page, err := h.pageService.LoadTemplate(path)
	if err != nil {
		log.Fatal(err)
	}

	err = page.Render(w)
	if err != nil {
		log.Fatal(err)
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
