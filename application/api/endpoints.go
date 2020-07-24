package api

import (
	"context"

	"github.com/MihaiBlebea/application/go-broadcast/hello"
	"github.com/go-kit/kit/endpoint"
)

// Endpoints is a data struct that holds the api endpoints
type Endpoints struct {
	HelloWorld endpoint.Endpoint
}

// MakeEndpoints returns the endpoints struct
func MakeEndpoints(helloService hello.Service) Endpoints {
	return Endpoints{
		HelloWorld: makeHelloWorldEndpoint(helloService),
	}
}

func makeHelloWorldEndpoint(service hello.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(HelloWorldRequest)
		msg, err := service.BroadcastMessage(req.Name, req.Age, req.Template)
		if err != nil {
			return HelloWorldResponse{Success: false}, err
		}

		return HelloWorldResponse{Message: msg, Success: true}, nil
	}
}
