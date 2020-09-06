package api

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/MihaiBlebea/list/lists"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sirupsen/logrus"
)

type server struct {
	list    lists.Service
	handler http.Handler
	logger  *logrus.Logger
}

// Server interface
type Server interface {
	Handler() *http.Handler
}

// New returns a new http server service
func New(list lists.Service, logger *logrus.Logger) Server {
	httpServer := server{
		list:   list,
		logger: logger,
	}

	r := mux.NewRouter()

	r.Methods("POST").Path("/lead").HandlerFunc(httpServer.PostLeadHandler)

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

func (s *server) Handler() *http.Handler {
	return &s.handler
}

func (s *server) PostLeadHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	reqID := s.requestID(8)

	s.logger.WithFields(logrus.Fields{
		"Url": path,
		"Id":  reqID,
	}).Info("Request started")
	start := time.Now()

	defer s.logger.WithFields(logrus.Fields{
		"Url":      path,
		"Id":       reqID,
		"Duration": time.Since(start),
	}).Info("Request ended")

	var req struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.list.AddContact(req.Name, req.Email)
	if err != nil {
		http.Error(w, err.Error(), 400)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(
		[]byte(`{"OK": "true"}`),
	)
}

func (s *server) requestID(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)

	for i := range b {
		b[i] = charset[random.Intn(len(charset))]
	}

	return string(b)
}
