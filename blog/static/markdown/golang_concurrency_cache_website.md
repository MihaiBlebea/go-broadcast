---
Title: Build a concurrent, concurrent server that can handle millions of requests with Golang
Summary: This is the second article
Image: https://thetravelvertical.com/wp-content/uploads/2019/11/dog-3822183_1920-1024x608.jpg
Tags:
    - golang
    - cache
    - concurrency
Slug: golang-concurrent-multithreaded-server-with-cache
# Published: 
---

## Step 1. Build a simple server with GO

To create the project, run this command:

- `mkdir dog_ceo && cd ./dog_ceo && go mod init github.com/<your-github-username>/dog-ceo && touch main.go`

This will create the folder, navigate into it, initialize the go module and create the `main.go` file.

When you navigate to your folder, you should see this files inside it:

- `go.mod` - This is simmilar to `package.json` in Node

- `go.sum` - This file locks the dependency versions. SOmething like `package.lock`

- `main.go` - This will be our entry file

We will keep the `main.go` file very light and use the package system to separate the logic of our website.

Let's install our first package, the `logrus` logger. Just run this command in the root filder:

- `go get github.com/sirupsen/logrus`

Let's write some code in the `main.go` file:

```go
package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
}
```

This will use the `logrus` package for logging errors in our website (we will also log a bunch of other stuff, but it's important to have the logs in a **nice json format**).

Now that we have the main file sorted, let's start working on the different packages that will support our website.

For a start, we will implement 3 packages:

- 1. dog package - will hold logic for fetching the images from the `dog.ceo` API

- 2. template package - will be responsable for rendering the static templates - you can think of this as the View in the MVC pattern)

- 3. api package - this will be a thin wrapper around the `gorilla/mux` router. It will hold the endpoints and the http handlers

I have chosen this order as every package depends on the previous one, so tackling them in this order would make the most sense.

The weird named `dog package` will have just 3 files `logic.go`, `service.go` and `logic.go`. I really like this way of managing the files inside a package, so I will continue using this pattern down this article.

- `model.go` file

```go
package dog

// Breed model
type Breed string

func toBreed(val string) Breed {
	return Breed(val)
}

// Dog model
type Dog string

func toDog(val string) Dog {
	return Dog(val)
}
```

This is very streight forward.

The `Breed` will be a custom type based on the string.

Same as `Dog`.

Both Breed and Dog will have methods for casting from a `string`, `toBreed()` and `toDog()`.

Notice that I kept this methods private so they will not be accesible from anywhere outside the `dog` package (we really have to do something about the package name).

- `service.go` file

```go
package dog

// Service interface
type Service interface {
	AllDogs() ([]Dog, error)
}
```
The service is only requesting one method to be implemented - `AllDogs()` that returns a slice of `Dog` and an `error`.

Nothing much to say here, so let's move to the service that implements this interface.

- `logic.go` file

```go
package dog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/sirupsen/logrus"
)

const url = "https://dog.ceo/api" // dog.ceo base api url

type service struct {
	url    string
	logger *logrus.Logger
}

// New retruns a new dog ceo service
func New(logger *logrus.Logger) Service {
	return &service{url, logger}
}

func (s *service) breeds() ([]Breed, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%s/breeds/list/all",
			s.url,
		),
		nil,
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []Breed{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data struct {
		Message map[string][]string `json:"message"`
		Status  string              `json:"status"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return []Breed{}, err
	}

	var breeds []Breed
	for breed := range data.Message {
		breeds = append(breeds, toBreed(breed))
	}

	sort.SliceStable(breeds, func(i, j int) bool {
		return breeds[i] < breeds[j]
	})

	return breeds, nil
}

func (s *service) dogs(breed Breed) ([]Dog, error) {
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"%s/breed/%v/images",
			s.url,
			breed,
		),
		nil,
	)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []Dog{}, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var data struct {
		Message []string `json:"message"`
		Status  string   `json:"status"`
	}

	if err := json.Unmarshal(body, &data); err != nil {
		return []Dog{}, err
	}

	var dogs []Dog
	for _, dog := range data.Message {
		dogs = append(dogs, toDog(dog))
	}

	return dogs, nil
}

func (s *service) AllDogs() ([]Dog, error) {
	breeds, err := s.breeds()
	if err != nil {
		return nil, err
	}

	type result struct {
		index int
		dogs  []Dog
		err   error
	}
	resultCh := make(chan result)

	for index, breed := range breeds {
		go func(index int, breed Breed) {
			d, err := s.dogs(breed)
			if err != nil {
				resultCh <- result{index: index, err: err}
			}
			resultCh <- result{index: index, dogs: d}
		}(index, breed)
	}

	var results []result
	for i := 0; i < len(breeds); i++ {
		results = append(results, <-resultCh)
	}

	// Sort the results by index
	sort.SliceStable(results, func(a, b int) bool {
		return results[a].index < results[b].index
	})

	var dogs []Dog
	for _, res := range results {
		if res.err == nil {
			dogs = append(dogs, res.dogs...)
		}
	}

	s.logger.WithFields(logrus.Fields{
		"dogs count":     len(dogs),
		"breeds count":   len(breeds),
		"requests count": len(breeds) + 1,
	}).Info("Dog request received")

	return dogs, nil
}
```

To simulate a time consumming process I have chosen to fetch the breeds of dogs from the api, then loop over the list and fetch all the dogs for each of the breeds.

I timed this reqeusts to about 6 seconds locally, so this should do it.

What is more interesting here is the `AllDogs` method - only exposed method outside of the package.

One refactoring that I did inside this method is the addition of a go routine, that can fetch the dogs by breed in a non blocking way. This really reduced the requests total time down to a couple of milliseconds.

Before starting to work on the server, we will build the template package.

Create a new folder `/template` in the root folder.

Let's split the package into:

- `service.go` - This file will hold the interface for our package service

- `logic.go` - This will be our main package file, it will contain our main service which implements the package service interface

- `model.go` - This is more or less self-explanatory

Let's take them one by one.

### Service.go

```go
package template

// Service - template service interface
type Service interface {
	Load(path string) (*Page, error)
}
```

Our main service interface will be pretty light. Just one method `Load()` which takes a string (path) and returns a pointer to a `Page` model and an error (hopefully not).

### Logic.go

```go
package template

import (
	"html/template"
	"io/ioutil"

	"github.com/MihaiBlebea/dog-ceo/dog"
)

type service struct {
	dogService dog.Service
}

// New returns a new template service
func New(dogService dog.Service) Service {
	return &service{dogService}
}

func (s *service) Load(path string) (*Page, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	tmp, err := template.New("Template").Parse(string(b))
	if err != nil {
		return nil, err
	}
	// fetch breeds and dogs
	dogs, err := s.dogService.AllDogs()
	if err != nil {
		return nil, err
	}

	return &Page{
		template: tmp,
		Dogs:     dogs,
	}, nil
}
```
The `service` struct has one single attribute that is the `dogService` which will be tasked with fetching the data from the third party API.

You can notice the implementation of the `Load` method that we defined on the interface.

The method takes the path of the template, reads the file and creates the actual template struct.

Then it uses the `dogService` to fetch all the images from the API.

In the end, it returns a `Page` struct that is composed from the template and the `dogs` slice.

### Model.go

```go
package template

import (
	"html/template"
	"io"
	"time"

	"github.com/MihaiBlebea/dog-ceo/dog"
)

// Page struct
type Page struct {
	template *template.Template
	Dogs     []dog.Dog
	Duration time.Duration // you can skip this
}

// Render the page
func (p *Page) Render(w io.Writer, duration time.Duration /* You can skip this */) error {
	p.Duration = duration

	err := p.template.Execute(w, p)
	if err != nil {
		return err
	}

	return nil
}

// DogsCount get dogs number by count
func (p *Page) DogsCount(count int) []dog.Dog {
	if len(p.Dogs) <= count {
		return p.Dogs
	}

	return p.Dogs[0:count]
}
```
We finally reached the model of our template package.

The `Page` is composed of a `html/template` (go standard library) and a slice of dogs (that is actaully not that bad as it sounds).

The model has just two methods:

- `Render` accepts a writer and a `time/duration` (this is optional) to render the page using the template. When we execute the template, we will pass the whole `Page` struct as params so we can access any method or attributes inside the template

- `DogsCount` we just defined this method so we can render a fix amount of dog images inside of the template



After this is completed, we will create our server with `gorilla/mux` package.

Just run this command in the root of your folder:

- `go get github.com/gorilla/mux`

You can check the `go.mod` file and see if both our dependencies were added to the list.

Your `go.mod` file should look like this:

```json
module github.com/<your-github-username>/dog-ceo

go 1.13

require (
	github.com/gorilla/mux v1.7.4
	github.com/sirupsen/logrus v1.6.0
)
```
There versions may be different depending on when you read this article.

Next, we will create the folder. Just add this file `http.go` in the `/api` folder (after you create the `api` folder in the root of your project).

The `/api` folder will act as a package to our project. Let's keep this minimal.

We will need just 2 routes to handle all the traffic hitting our website:

- `/` this will be the main route and will display the images in our gallery

- `/static` this will be the endpoint that we will call to get the html template and the css file (this is optional, as you can use a bundle the template binary into the GO app binary, but for this project we will keep them as static files)

This is how the `/api/http.go` file will look:

```go
package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path"
	"time"

	"github.com/MihaiBlebea/dog-ceo/template"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type server struct {
	templateService template.Service
	handler         http.Handler
	logger          *logrus.Logger
}

// Server interface
type Server interface {
	Handler() *http.Handler
}

// New returns a new http server service
func New(templateService template.Service, logger *logrus.Logger) Server {
	httpServer := server{
		templateService: templateService,
		logger:          logger,
	}

	r := mux.NewRouter()

	r.Methods("GET").Path("/").HandlerFunc(httpServer.indexHandler)

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

	httpServer.handler = r

	return &httpServer
}

func (h *server) Handler() *http.Handler {
	return &h.handler
}

func (h *server) staticFolderPath() string {
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

func (h *server) indexHandler(w http.ResponseWriter, r *http.Request) {
	h.logger.WithFields(logrus.Fields{
		"url": r.URL.String(),
	}).Info("HTTP request started")
	start := time.Now()

	defer h.logger.WithFields(logrus.Fields{
		"duration": time.Since(start).Nanoseconds(),
	}).Info("HTTP request ended")

	path := h.staticFolderPath() + "html/index.gohtml"
	page, err := h.templateService.Load(path)

	if err != nil {
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
	}

	page.Render(w, time.Since(start))
}

```

## Step 2. Fetch the data from the third-party API

## Step 3. Display the data with server-side rendered GO templates

## Step 4. Refactoring - Add cache to improve speed

## Step 5. Refactoring - Add worker to cache the data in the background

## Step 6. Use docker to bundle build and deploy the GO server
