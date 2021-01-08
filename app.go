package sageengine

// App holds the information about an app which can be used in tasks.
type App struct {
	Version  uint
	Metadata struct {
		Name    string
		Version string
	}
	Auths []struct {
		Kind   AuthKind
		Fields []struct {
			Key              string
			Label            string
			Tip              string
			IsRequired       bool      `yaml:"is_required" json:"is_required"`
			IsAdministrative bool      `yaml:"is_administrative" json:"is_administrative"`
			InputKind        InputKind `yaml:"input_kind" json:"input_kind"`
			DefaultValue     string    `yaml:"default_value" json:"default_value"`
			Dropdown         struct {
				Kind           DropdownKind
				AllowsMultiple bool `yaml:"allows_multiple" json:"allows_multiple"`
				AllowsCustom   bool `yaml:"allows_custom" json:"allows_custom"`
				Options        []string
			}
		}
		API API
	}
	Operations []struct {
		Kind             OperationKind
		Key              string
		Label            string
		Tip              string
		IsRequired       bool      `yaml:"is_required" json:"is_required"`
		IsWriteOp        bool      `yaml:"is_write_op" json:"is_write_op"`
		IsIdentification bool      `yaml:"is_identification" json:"is_identification"`
		ResourceNoun     string    `yaml:"resource_noun" json:"resource_noun"`
		InputKind        InputKind `yaml:"input_kind" json:"input_kind"`
		Dropdown         struct {
			Kind           DropdownKind
			AllowsMultiple bool `yaml:"allows_multiple" json:"allows_multiple"`
			AllowsCustom   bool `yaml:"allows_custom" json:"allows_custom"`
			Options        []string
		}
		API API
	}
}

// API specifies how a resource is resoved, fetched, updated, etc.
type API struct {
	Code     string
	Language string
	Form     struct {
		Method  string
		URL     string
		Headers string
		Body    string
	}
}

// AuthKind is the type of authorisation an App supports.
type AuthKind string

// The different types of authorization.
const (
	OAUTH2 AuthKind = "OAUTH2"
)

// InputKind is the type of user input.
type InputKind string

// The different types of user input.
const (
	EMAIL  InputKind = "EMAIL"
	SELECT InputKind = "SELECT"
)

// DropdownKind is the type of dropdown.
type DropdownKind string

// The different types of dropdown.
const (
	STATIC  DropdownKind = "STATIC"
	DYNAMIC DropdownKind = "DYNAMIC"
)

// OperationKind is the type of operation.
type OperationKind string

// The different types of operation.
const (
	TRIGGER OperationKind = "TRIGGER"
	ACTION  OperationKind = "ACTION"
	SEARCH  OperationKind = "SEARCH"
)
