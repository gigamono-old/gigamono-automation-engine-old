package mainserver

import (
	"context"
	"fmt"
	"net"

	"github.com/gigamono/gigamono/pkg/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (server *MainServer) grpcServe(listener net.Listener) error {
	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	generated.RegisterWorkflowMainServerServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello replies with message.
func (server *MainServer) SayHello(ctx context.Context, msg *generated.Message) (*generated.Message, error) {
	engineMsg := "Main Server replies: " + msg.Content
	fmt.Println(engineMsg)
	response := generated.Message{
		Content: engineMsg,
	}
	return &response, nil
}
