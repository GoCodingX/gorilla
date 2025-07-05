package repository

import (
	"time"

	"github.com/google/uuid"
)

type GetQuotesParams struct {
	Author          *string
	CursorCreatedAt *time.Time
	CursorID        *uuid.UUID
}

type QuotesCursor struct {
	CreatedAt time.Time
	ID        uuid.UUID
}
