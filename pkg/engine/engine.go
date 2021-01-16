package engine

import (
	"fmt"
	"net"
	"reflect"
	"strings"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"google.golang.org/grpc"

	"github.com/sageflow/sagedb/pkg/database"
	"github.com/sageflow/sagedb/pkg/models"
	"github.com/sageflow/sageutils/pkg/configs"
)

// Engine represents an engine instance.
// Used as a gRPC server, Engine expects an API server, an API DB and an Auth server.
// Used as an API, Engine expects an API DB. API and Auth servers are optional.
// If there is no authorisation with the Auth server, the API does not require a Context as well.
type Engine struct {
	ID      uuid.UUID
	AppPool []App
	Port    string
}

// NewEngine creates a new workflow engine.
func NewEngine(db *database.DB) Engine {
	// Create engine in the database.
	engine := models.CreateEngine(db)
	return Engine{
		ID: engine.ID,
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

	// Create a gRPC server that uses TCP listener.
	grpcServer := grpc.NewServer()
	if err := grpcServer.Serve(listener); err != nil {
		return err
	}

	// Registers a gRPC client.

	// Accept ExecuteWorkflow call request from clients.

	// Execute(workflow Workflow, context Context) // engine.proto

	// If it happens.

	// Gets json from the db

	// Call engine.ExecuteWorkflowJSON

	return nil
}

// ExecuteWorkflow takes a workflow object and executes it.
func (engine *Engine) ExecuteWorkflow(workflow *Workflow, context *Context) error {
	// Starts workflow on a separate goroutine
	go workflow.Execute(context)

	// Wait workflow response. Return any error to user

	// Runs workflow again based on specific time interval.

	return nil
}

// ExecuteWorkflowString submits a workflow string in the following format (JSON, TOML, YAML) for execution.
func (engine *Engine) ExecuteWorkflowString(format configs.ConfigFormat, str string, context *Context) error {
	app := App{}
	reader := strings.NewReader(str)

	// Set viper to parse format.
	viper.SetConfigType(format.String())
	viper.ReadConfig(reader)

	// Convert format into Workflow object.
	if err := viper.Unmarshal(&app, getCustomDecoder()); err != nil {
		return err
	}

	fmt.Println("\n=========== " + format.String() + " =============")
	fmt.Println("\nWorkflow object =", app)

	return engine.ExecuteWorkflow(&Workflow{}, context)
}

func getCustomDecoder() viper.DecoderConfigOption {
	return viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(
		func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
			switch t {
			case reflect.TypeOf(UUID{}):
				parsedID, err := uuid.Parse(data.(string))
				if err != nil {
					return UUID{}, err
				}
				id := UUID(parsedID)
				return id, nil
			}
			return data, nil
		},
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
	))
}
