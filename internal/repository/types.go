package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            uuid.UUID `bun:",pk,type:uuid,notnull"`
	Username      string    `bun:",notnull"`
	CreatedAt     time.Time `bun:",notnull"`
	UpdatedAt     time.Time `bun:",notnull"`
}

type Author struct {
	bun.BaseModel `bun:"table:authors,alias:a"`
	ID            uuid.UUID `bun:",pk,type:uuid,notnull"`
	Name          string    `bun:",notnull"`
	CreatedAt     time.Time `bun:",notnull"`
	UpdatedAt     time.Time `bun:",notnull"`
}

type Quote struct {
	bun.BaseModel `bun:"table:quotes,alias:q"`
	ID            uuid.UUID `bun:",pk,type:uuid,notnull"`
	Text          string    `bun:",notnull"`
	*Author       `bun:",notnull,rel:belongs-to,join:author_id=id"`
	*User         `bun:",notnull,rel:belongs-to,join:creator_user_id=id"`
	CreatedAt     time.Time `bun:",notnull"`
	UpdatedAt     time.Time `bun:",notnull"`
}
