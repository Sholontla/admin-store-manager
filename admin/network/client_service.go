package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/valyala/fasthttp"
)

func main() {
	url := "http://localhost:8080/get-config"

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	req.SetRequestURI(url)
	req.Header.Set("Authorization", "Bearer client1:a7d74b46d3e7f14a32b6c61fcb1e28d4528c44175b54af2b1037fcdefa0a1121")

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if err := fasthttp.Do(req, resp); err != nil {
		log.Fatalf("Error getting configuration data: %v", err)
	}

	bodyBytes := resp.Body()
	response := Response{}
	err := json.Unmarshal(bodyBytes, &response)
	if err != nil {
		log.Fatalf("Error parsing response: %v", err)
	}

	fmt.Printf("Received configuration data: %v\n", response)
}

type Response struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}
