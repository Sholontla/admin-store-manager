package main

import (
	"fmt"
	"log"

	pb "service/stores/case1/internal/proto"

	"github.com/zeromq/goczmq"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Create a ZeroMQ publisher socket
	publisher, err := goczmq.NewPub("tcp://*:5555")
	if err != nil {
		log.Fatal(err)
	}
	defer publisher.Destroy()

	// Create a channel to receive messages from the client
	msgChan := make(chan []byte)

	// Start a goroutine to read from the channel and publish messages
	go func() {
		var msgs [][]byte
		for {
			select {
			case msg := <-msgChan:
				msgs = append(msgs, msg)
			default:
				if len(msgs) > 0 {
					if err := publisher.SendFrame(msgs[0], 0); err != nil {
						log.Println(err)
					}
					msgs = nil // reset msgs slice
				}
			}
		}
	}()

	// Create a message to send to the client
	myMessage := &pb.MyMessage{
		Id:    "myId",
		Value: 123,
		NestedMessages: []*pb.MyNestedMessage{
			{Name: "name1", Score: 1.2},
			{Name: "name2", Score: 3.4},
		},
	}

	// Serialize the message
	serializedMsg, err := proto.Marshal(myMessage)
	if err != nil {
		log.Fatal(err)
	}

	// Send the message to the client
	msgChan <- serializedMsg

	// Wait for input to exit
	fmt.Println("Server running...")
	fmt.Println("Press any key to exit")
	fmt.Scanln()
	fmt.Println("Exiting...")
}
