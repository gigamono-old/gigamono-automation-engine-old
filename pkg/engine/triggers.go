package engine

// Triggerer represents a workflow trigger.
type Triggerer interface {
	RegisterWorkflowInstance(instance *WorkflowInstance)
	Watch()
}


// func NewTrigger(workflow *Workflow) Triggerer {
// }
