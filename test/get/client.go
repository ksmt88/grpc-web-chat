package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	chat "github.com/ksmt88/grpc-web-chat/proto"
)

func main() {
	fmt.Println("Hello Client!")

	cc, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure(), grpc.WithKeepaliveParams(
		keepalive.ClientParameters{
			Time:                0,
			Timeout:             10 * time.Minute,
			PermitWithoutStream: true,
		},
	))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer cc.Close()

	c := chat.NewChatClient(cc)

	getMessage(c)
}

func getMessage(c chat.ChatClient) {
	fmt.Println("Starting")

	messageStream, err := c.GetMessages(context.Background(), &empty.Empty{})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	for {
		msg, err := messageStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		log.Printf("Response from GetMessages: name[%v], message[%v], createdAt[%v].", msg.GetName(), msg.GetMessage(), msg.GetCreatedAt())
	}
}
