package pg

import (
	"github.com/uptrace/bun"
)

type Repository struct {
	client *bun.DB
}

func NewRepository(client *bun.DB) *Repository {
	return &Repository{
		client: client,
	}
}
