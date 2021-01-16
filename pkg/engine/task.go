package engine

// Task is an executable step in a workflow.
type Task struct {
	Kind             TaskKind
	Name             string
	Index            uint
	Dependencies     []uint
	ExecutionContext ExecutionContext `mapstructure:"execution_context"`
	AppID            UUID             `mapstructure:"app_id"`
	Fields           map[string][]string
	AppName          string `mapstructure:"app_name"`
	AccountID        UUID   `mapstructure:"account_id"`
}
