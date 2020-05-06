package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"google.golang.org/grpc"

	chat "github.com/ksmt88/grpc-web-chat/proto"
)

func main() {
	fmt.Println("Hello Client!")

	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := chat.NewChatClient(cc)

	postMessage(c, "a", "message from a.")
	time.Sleep(time.Second)
	postMessage(c, "b", "message from b.")
	time.Sleep(time.Second)
	postMessage(c, "c", "message from c.")
}

func postMessage(c chat.ChatClient, name string, message string) {

	req := &chat.Message{
		Name:      name,
		Message:   message,
		CreatedAt: ptypes.TimestampNow(),
	}
	res, err := c.PostMessage(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PostMessage RPC: %v", err)
	}
	log.Printf("Response from PostMessage: %v", res)
}
