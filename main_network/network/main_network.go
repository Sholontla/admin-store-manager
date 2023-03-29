package network

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net"
	"strings"

	"github.com/valyala/fasthttp"
)

type Response struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

func StartServer() {
	authorizedClients := map[string]string{
		"client1": "a7d74b46d3e7f14a32b6c61fcb1e28d4528c44175b54af2b1037fcdefa0a1121", // replace with your own hashed password/key
		"client2": "f73ad37c4b236078b3f3ed4d4b51026b36f7bb2c2d35675d8ce28f03fa49edf6",
	}

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		if string(ctx.Path()) != "/get-config" {
			ctx.Error("Unsupported path", fasthttp.StatusNotFound)
			return
		}

		authHeader := string(ctx.Request.Header.Peek("Authorization"))
		authParts := strings.Split(authHeader, " ")

		if len(authParts) != 2 || authParts[0] != "Bearer" {
			ctx.Error("Invalid authorization header", fasthttp.StatusUnauthorized)
			return
		}

		clientID := authParts[1]

		if password, ok := authorizedClients[clientID]; !ok || password != hashPassword("my_password") {
			ctx.Error("Invalid client credentials", fasthttp.StatusUnauthorized)
			return
		}

		response := Response{
			Message:     "Configuration data",
			Description: "This is the data generated by the server based on the client's request",
		}

		jsonResponse, err := json.Marshal(response)
		if err != nil {
			ctx.Error("Error generating response", fasthttp.StatusInternalServerError)
			return
		}

		ctx.Response.Header.SetContentType("application/json")
		ctx.Response.SetBody(jsonResponse)
	}

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()

	log.Println("Listening on port 8080...")

	fasthttp.Serve(listener, requestHandler)
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
