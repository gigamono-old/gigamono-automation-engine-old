package engine

import (
	"github.com/google/uuid"
)

// Context holds information about user, session, etc.
type Context struct {
	UserID  uuid.UUID
	AppPool *[]App
}
