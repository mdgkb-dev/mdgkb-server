package mocks

import (
	"mdgkb/mdgkb-server/models"

	"github.com/google/uuid"
)

var MockFull = models.PathPermission{
	ID:         uuid.NullUUID{UUID: uuid.New(), Valid: true},
	Resource:   "2",
	GuestAllow: true,
}
