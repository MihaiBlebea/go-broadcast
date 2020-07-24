package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MihaiBlebea/application/go-broadcast/api"
	"github.com/MihaiBlebea/application/go-broadcast/hello"
)

func main() {
	service := hello.New()

	endpoints := api.MakeEndpoints(service)

	handler := api.NewHTTPServer(context.Background(), endpoints)

	httpPort := fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))
	log.Println("Application started HTTP server on port " + httpPort)
	err := http.ListenAndServe(httpPort, handler)
	if err != nil {
		log.Fatal(err)
	}
}
