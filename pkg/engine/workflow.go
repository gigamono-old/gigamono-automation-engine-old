package engine

// Workflow represents a runnable workflow.
type Workflow struct {
	ID       UUID
	Version  uint
	Metadata struct {
		Name              string
		ExecutionContexts []ExecutionContext `mapstructure:"execution_contexts"`
	}
	Tasks []Task
}

// Execute starts the execution of workflow run.
func (workflow *Workflow) Execute(context *Context) error {
	// TODO
	return nil
}

func (task *Task) execute(workflow *Workflow, context *Context) error {
	// TODO
	return nil
}
