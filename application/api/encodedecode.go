package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeHelloWorldRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var req HelloWorldRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}

	return req, nil
}
