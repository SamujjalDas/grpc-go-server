package dummy

import (
	"github.com/gofrs/uuid"
	"github.com/google/uuid"
)

type Dummy struct {
	UserId   uuid.UUID
	UserName string
}
