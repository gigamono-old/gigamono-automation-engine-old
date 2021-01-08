package sageengine

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"gopkg.in/yaml.v3"

	"github.com/sageflow/sagedb"
	"github.com/sageflow/sagedb/models"
)

// Context holds information about user, session, etc.
type Context struct {
	UserID  uuid.UUID
	AppPool *[]App
}

// Engine represents an engine instance.
// Used as a gRPC server, Engine expects an API server, an API DB and an Auth server.
// Used as an API, Engine expects an API DB. API and Auth servers are optional.
// If there is no authorisation with the Auth server, the API does not require a Context as well.
type Engine struct {
	ID      uuid.UUID
	AppPool []App
}

// ExecutionContext refers to how the engine is going to run each task.
type ExecutionContext string

// The different types of execution contexts.
const (
	PROTECTED ExecutionContext = "PROTECTED" // Sandboxed code execution
	BARE      ExecutionContext = "BARE"      // Non-sandboxed code execution
)

// UUID aliases Google's UUID type to allow custom unmarhsalling.
type UUID uuid.UUID

// UnmarshalJSON is a custom unmarshaller for UUID.
func (id *UUID) UnmarshalJSON(bytes []byte) error {
	newID, err := uuid.ParseBytes(bytes)
	if err != nil {
		return err
	}
	*id = UUID(newID)
	return nil
}

// UnmarshalYAML is a custom unmarshaller for UUID.
func (id *UUID) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var corpus string
	if err := unmarshal(&corpus); err != nil {
		return err
	}

	parsedID, err := uuid.Parse(corpus)
	if err != nil {
		return err
	}

	*id = UUID(parsedID)

	return nil
}

// NewEngine creates a new workflow engine.
func NewEngine(db *sagedb.DB) Engine {
	// Create engine in the database.
	engine := models.CreateEngine(db)
	return Engine{
		ID: engine.ID,
	}
}

// Start starts the engines gRPC server.
func (engine *Engine) Start(port string) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	// Registers a gRPC client

	// TODO
	// Waits for client to ask for ExecuteWorkflow call
	// Execute(workflow Workflow, context Context) // engine.proto
	// If it happens.
	// Gets json from the db
	// Call engine.ExecuteWorkflowJSON
	return nil
}

// ExecuteWorkflowYAML submits a workflow in yaml for execution.
func (engine *Engine) ExecuteWorkflowYAML(yamlString string, context Context) error {
	// Convert YAML into Workflow object.
	workflow := Workflow{}
	err := yaml.Unmarshal([]byte(yamlString), &workflow)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\nWorkflow object =", workflow)
	fmt.Println("\nWorkflow.Tasks[0].AppID =", workflow.Tasks[0].AppID)

	// Starts the workflow immediately on a separate goroutine
	// go workflow.Execute(context Context)

	// TODO
	// Wait for exit response.
	// Returns any error to user. Fail fast
	// Runs workflow again based on specific time interval.
	return nil
}

// ExecuteWorkflowJSON submits a workflow in JSON for execution.
func (engine *Engine) ExecuteWorkflowJSON(yamlString string, context Context) error {
	// Convert JSON into Workflow object.
	workflow := Workflow{}
	err := json.Unmarshal([]byte(yamlString), &workflow)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("\nWorkflow object =", workflow)
	fmt.Println("\nWorkflow.Tasks[0].AppID =", workflow.Tasks[0].AppID)

	// Starts the workflow immediately on a separate goroutine
	// go workflow.Execute(context Context)

	// TODO
	// Wait for exit response.
	// Returns any error to user. Fail fast
	// Runs workflow again based on specific time interval.
	return nil
}
