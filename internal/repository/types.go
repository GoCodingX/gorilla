package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Quote struct {
	bun.BaseModel `bun:"table:quotes,alias:q"`
	ID            uuid.UUID `bun:",pk,type:uuid,notnull"`
	Text          string
	AuthorId      string
	CreatorId     string
	CreatedAt     time.Time `bun:",notnull"`
	UpdatedAt     time.Time `bun:",notnull"`
}
