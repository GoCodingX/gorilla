package pg

import (
	"context"
	"fmt"

	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/pkg/db"
)

func (r *Repository) CreateAuthor(ctx context.Context, author *repository.Author) error {
	_, err := r.db.NewInsert().Model(author).Exec(ctx)
	if err != nil {
		if db.IsUniqueViolation(err) {
			return repository.NewAlreadyExistsError(author.Name, err)
		}

		return fmt.Errorf("failed to persist author: %w", err)
	}

	return nil
}
