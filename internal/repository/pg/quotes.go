package pg

import (
	"context"
	"fmt"

	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/pkg/db"
)

func (r *Repository) CreateQuote(ctx context.Context, quote *repository.Quote) error {
	_, err := r.db.NewInsert().Model(quote).Exec(ctx)
	if err != nil {
		if db.IsForeignKeyViolation(err) {
			return repository.NewInvalidReferenceError(quote.AuthorID.String(), err)
		}

		return fmt.Errorf("failed to persist quote: %w", err)
	}

	return nil
}

func (r *Repository) GetQuotes(ctx context.Context, params *repository.GetQuotesParams) ([]repository.Quote, error) {
	var quotes []repository.Quote

	query := r.db.NewSelect().
		Model(&quotes).
		Relation("Author").
		Order("q.created_at DESC").
		Limit(10)

	if params.Author != nil && len(*params.Author) > 0 {
		query = query.Where("author.name = ?", params.Author)
	}

	if err := query.Scan(ctx); err != nil {
		return nil, err
	}

	return quotes, nil
}
