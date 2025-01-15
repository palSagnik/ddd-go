package valueobjects

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	from      uuid.UUID
	to        uuid.UUID
	amt       int
	createdAt time.Time
}
