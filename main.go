package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/go-redis/redis/v7"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ksmt88/grpc-web-chat/proto"
)

const Channel = "chatChannel"

type server struct{}

func (s *server) GetMessages(request *empty.Empty, stream chat.Chat_GetMessagesServer) error {
	client := NewRedisClient()

	pubsub := client.Subscribe(Channel)
	defer pubsub.Close()

	_, err := pubsub.Receive()
	if err != nil {
		log.Fatal(err)
	}

	ch := pubsub.Channel()

	// Consume messages.
	for msg := range ch {
		var message chat.Message
		err := json.Unmarshal([]byte(msg.Payload), &message)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(msg.Channel, msg.Payload)

		if err := stream.Send(&message); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) PostMessage(ctx context.Context, request *chat.Message) (*chat.Result, error) {
	client := NewRedisClient()

	if request.GetCreatedAt() == nil {
		request.CreatedAt = ptypes.TimestampNow()
	}

	message, _ := json.Marshal(request)

	_ = client.Publish(Channel, message).Err()

	result := &chat.Result{
		Result: true,
	}

	return result, nil
}

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}

func main() {
	log.Println("starting server")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	chat.RegisterChatServer(s, &server{})

	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
