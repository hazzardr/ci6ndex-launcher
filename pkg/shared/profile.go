package shared

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID          uint
	Name        string
	Description string
	ModUUIDs    []uuid.UUID
	CreatedAt   time.Time
}
