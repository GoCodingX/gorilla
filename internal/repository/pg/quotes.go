package pg

import (
	"context"
	"fmt"

	"github.com/GoCodingX/gorilla/internal/repository"
	"github.com/GoCodingX/gorilla/pkg/db"
)

const defaultLimit = 10

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

func (r *Repository) GetQuotes(ctx context.Context, params *repository.GetQuotesParams) ([]repository.Quote, *repository.QuotesCursor, error) {
	var quotes []repository.Quote

	query := r.db.NewSelect().
		Model(&quotes).
		Relation("Author").
		Order("q.created_at DESC").
		Order("q.id DESC").
		// request an extra record to determine whether there's a next page
		Limit(defaultLimit + 1)

	if params.Author != nil && len(*params.Author) > 0 {
		query = query.Where("author.name = ?", params.Author)
	}

	if params.Author != nil && len(*params.Author) > 0 {
		query = query.Where("author.name = ?", *params.Author)
	}

	if params.CursorCreatedAt != nil && params.CursorID != nil {
		query = query.Where("(q.created_at, q.id) < (?, ?)", *params.CursorCreatedAt, *params.CursorID)
	}

	if err := query.Scan(ctx); err != nil {
		return nil, nil, err
	}

	var nextCursor *repository.QuotesCursor

	if len(quotes) > defaultLimit {
		last := quotes[defaultLimit-1]

		nextCursor = &repository.QuotesCursor{
			CreatedAt: last.CreatedAt,
			ID:        last.ID,
		}

		// drop extra record
		quotes = quotes[:defaultLimit]
	}

	return quotes, nextCursor, nil
}
