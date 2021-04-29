package resolver

import "github.com/gigamono/gigamono/pkg/inits"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver holds dependencies like App.
type Resolver struct {
	*inits.App
}
