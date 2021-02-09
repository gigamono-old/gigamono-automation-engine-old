package server

import (
	"context"
	"fmt"
	"net"

	"github.com/sageflow/sageengine/internal/proto"
	"github.com/sageflow/sageengine/pkg/engine"

	"github.com/sageflow/sageflow/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// EngineServer is a grpc server with an engine.
type EngineServer struct {
	Port   string
	Engine engine.Engine
}

// NewEngineServer creates a new server instance.
func NewEngineServer(db *database.DB) EngineServer {
	return EngineServer{
		Engine: engine.NewEngine(db),
	}
}

// Listen starts a new gRPC server that listens on specified port.
func (server *EngineServer) Listen(port string) error {
	server.Port = port // Set port.

	// Listen on port using TCP.
	listener, err := net.Listen("tcp", ":"+server.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	proto.RegisterEngineServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello says Hello
func (server *EngineServer) SayHello(ctx context.Context, msg *proto.Message) (*proto.Message, error) {
	engineMsg := "Engine replies: " + msg.Content
	fmt.Println(engineMsg)
	response := proto.Message{
		Content: engineMsg,
	}
	return &response, nil
}
