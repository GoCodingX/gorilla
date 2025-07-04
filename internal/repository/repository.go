//go:generate go run go.uber.org/mock/mockgen -package=repositorytest -source=repository.go -destination=repositorytest/repository.go .

package repository

import (
	"context"
)

type Repository interface {
	CreateQuote(ctx context.Context, quote *Quote) error
	GetQuotes(ctx context.Context, params *GetQuotesParams) ([]Quote, error)
	CreateAuthor(ctx context.Context, author *Author) error
}
