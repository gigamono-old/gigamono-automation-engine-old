package runnablesupervisor

import (
	"context"
	"fmt"
	"net"

	"github.com/gigamono/gigamono/pkg/services/proto/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func (supervisor *RunnableSupervisor) grpcServe() error {
	listener, err := net.Listen(
		"tcp",
		fmt.Sprint(
			":",
			supervisor.Config.Services.Types.WorkflowEngine.PrivatePorts.RunnableSupervisor,
		),
	)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer() // Create a gRPC service.

	// Register gRPC service.
	generated.RegisterWorkflowRunnableSupervisorServer(grpcServer, supervisor)
	reflection.Register(grpcServer)

	return grpcServer.Serve(listener) // Listen for requests.
}

// SayHello replies with message.
func (supervisor *RunnableSupervisor) SayHello(ctx context.Context, msg *generated.Message) (*generated.Message, error) {
	engineMsg := "Runnable Supervisor replies: " + msg.Content
	fmt.Println(engineMsg)
	response := generated.Message{
		Content: engineMsg,
	}
	return &response, nil
}
