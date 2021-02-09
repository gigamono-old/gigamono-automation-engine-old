package engine

import (
	"context"
	"fmt"
	"net"

	"github.com/sageflow/sageengine/internal/proto"

	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/sageflow/sageflow/pkg/database"
	"github.com/sageflow/sageflow/pkg/database/models"
)

// Engine represents an engine instance.
// Used as a gRPC server, Engine expects an API server, an API DB and an Auth server.
// Used as an API, Engine expects an API DB. API and Auth servers are optional.
// If there is no authorisation with the Auth server, the API does not require a Context as well.
type Engine struct {
	Model *models.Engine
	Port  string
	DB    *database.DB
}

// NewEngine creates a new workflow engine.
func NewEngine(db *database.DB) Engine {
	// Create engine in the database.
	engine := db.CreateEngine()
	return Engine{
		Model: engine,
	}
}

// Listen starts a new gRPC server that listens on specified port.
func (engine *Engine) Listen(port string) error {
	engine.Port = port // Set port.

	// Listen on port using TCP.
	listener, err := net.Listen("tcp", ":"+engine.Port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() // Create a gRPC server.

	// Register gRPC service.
	proto.RegisterEngineServiceServer(grpcServer, engine)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello says Hello
func (engine *Engine) SayHello(ctx context.Context, msg *proto.Message) (*proto.Message, error) {
	engineMsg := "Engine replies: " + msg.Content
	fmt.Println(engineMsg)
	response := proto.Message{
		Content: engineMsg,
	}
	return &response, nil
}

// ExecuteWorkflow takes a workflow object and executes it.
func (engine *Engine) ExecuteWorkflow(workflowID uuid.UUID) error {
	// If it exists.
	// Execute workflow
	return nil
}
