package server

import (
	"context"
	"fmt"
	"net"

	"github.com/gigamono/gigamono-workflow-engine/pkg/engine"
	"github.com/gigamono/gigamono/pkg/services/proto/generated"

	"github.com/gigamono/gigamono/pkg/inits"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// WorkflowEngineServer is a grpc server with an engine.
type WorkflowEngineServer struct {
	inits.App
	Engine engine.WorkflowEngine
}

// NewWorkflowEngineServer creates a new server instance.
func NewWorkflowEngineServer(app inits.App) (WorkflowEngineServer, error) {
	eng, err := engine.NewWorkflowEngine(&app)
	if err != nil {
		return WorkflowEngineServer{}, err
	}
	return WorkflowEngineServer{
		App:    app,
		Engine: eng,
	}, nil
}

// Listen starts a new gRPC server that listens on specified port.
func (server *WorkflowEngineServer) Listen() error {
	// Listen on port using TCP.
	listener, err := net.Listen("tcp", fmt.Sprint(":", server.Config.Services.Types.WorkflowEngine.Port))
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	generated.RegisterWorkflowEngineServiceServer(grpcServer, server)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello says Hello
func (server *WorkflowEngineServer) SayHello(ctx context.Context, msg *generated.Message) (*generated.Message, error) {
	engineMsg := "Engine replies: " + msg.Content
	fmt.Println(engineMsg)
	response := generated.Message{
		Content: engineMsg,
	}
	return &response, nil
}
