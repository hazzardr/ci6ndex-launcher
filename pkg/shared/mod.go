package shared

import "github.com/google/uuid"

type Mod struct {
	UUID       uuid.UUID
	Name       string
	Version    string
	Author     string
	Source     string
	Path       string
	WorkshopID uint
}
