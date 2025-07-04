package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Author struct {
	bun.BaseModel `bun:"table:authors,alias:a"`
	ID            uuid.UUID `bun:",pk,type:uuid,notnull"`
	Name          string    `bun:",notnull"`
	CreatedAt     time.Time `bun:",notnull"`
	UpdatedAt     time.Time `bun:",notnull"`
}

type Quote struct {
	bun.BaseModel   `bun:"table:quotes,alias:q"`
	ID              uuid.UUID `bun:",pk,type:uuid,notnull"`
	Text            string    `bun:",notnull"`
	AuthorID        uuid.UUID
	Author          *Author   `bun:",rel:belongs-to,join:author_id=id"`
	CreatorUsername string    `bun:",notnull"`
	CreatedAt       time.Time `bun:",notnull"`
	UpdatedAt       time.Time `bun:",notnull"`
}
