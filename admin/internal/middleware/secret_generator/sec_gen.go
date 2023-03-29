package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	key := make([]byte, 64)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	secretKey := base64.StdEncoding.EncodeToString(key)
	fmt.Println(secretKey)
}
