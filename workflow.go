package sageengine

// Workflow represents a runnable workflow.
type Workflow struct {
	Version  uint
	Metadata struct {
		Name              string
		ExecutionContexts []ExecutionContext `yaml:"execution_contexts" json:"execution_contexts"`
	}
	Tasks []Task
}

// Task is an executable step in a workflow.
type Task struct {
	Kind             OperationKind
	Name             string
	Index            uint
	Dependencies     []uint
	ExecutionContext ExecutionContext `yaml:"execution_context" json:"execution_context"`
	Fields           map[string]string
	AppName          string `yaml:"app_name" json:"app_name"`
	AppID            UUID   `yaml:"app_id" json:"app_id"`
	AccountID        UUID   `yaml:"account_id" json:"account_id"`
}

// Execute starts the execution of workflow run.
func (workflow *Workflow) Execute(context Context) error {
	// TODO
	return nil
}

func (task *Task) execute(workflow *Workflow, context Context) error {
	// TODO
	return nil
}
