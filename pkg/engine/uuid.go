package engine

import (
	"github.com/google/uuid"
)

// UUID aliases Google's UUID type to allow custom unmarhsalling.
type UUID uuid.UUID
