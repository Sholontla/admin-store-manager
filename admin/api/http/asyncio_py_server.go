package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
)

func AsyncConnServer() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	message := map[string]string{"endpoint": "t", "key1": "value1"}
	jsonData, _ := json.Marshal(message)

	// Send message to server
	fmt.Fprintf(conn, string(jsonData)+"\n")

	// Read response from server
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		response := scanner.Text()
		fmt.Println("Received response from server:", response)
		break
	}
}
