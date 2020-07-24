package api

// HelloWorldRequest request contract
type HelloWorldRequest struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Template string `json:"template"`
}

// HelloWorldResponse response contract
type HelloWorldResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
